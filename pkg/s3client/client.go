package s3client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	cfg "github.com/sojebsikder/go-boilerplate/internal/config"
)

type S3Client struct {
	client  *s3.Client
	presign *s3.PresignClient
	config  *cfg.Config
	bucket  string
}

func NewS3Client(cfg *cfg.Config) *S3Client {
	s3Cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.S3.AWSAccessKeyID,
				cfg.S3.AWSSecretAccessKey,
				"",
			),
		),
		config.WithRegion(cfg.S3.AWSRegion),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...any) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:               cfg.S3.AWSEndpoint,
						HostnameImmutable: true,
					}, nil
				},
			),
		),
	)

	if err != nil {
		panic(err)
	}

	s3Client := s3.NewFromConfig(s3Cfg, func(o *s3.Options) {
		o.UsePathStyle = true // IMPORTANT for MinIO
	})

	presignClient := s3.NewPresignClient(s3Client)

	return &S3Client{
		client:  s3Client,
		presign: presignClient,
		bucket:  cfg.S3.AWSBucket,
		config:  cfg,
	}
}

func (s *S3Client) SetBucket(bucket string) {
	s.bucket = bucket
}

func (s *S3Client) GetBucket() string {
	return s.bucket
}

// Upload file to S3
func (s *S3Client) UploadFile(ctx context.Context, file multipart.File, fileName string) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fileName),
		Body:   file,
		ACL:    "public-read", // optional
	})
	return err
}

func (s *S3Client) UploadBytes(
	ctx context.Context,
	data []byte,
	key string,
	contentType string, // optional but recommended
) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
		ACL:         "public-read", // remove if you don’t want public access
	})

	return err
}

func (s *S3Client) UploadBytesStream(
	ctx context.Context,
	data io.Reader,
	key string,
	contentType *string,
) error {
	var body io.Reader = data

	if contentType == nil {
		ct, newReader, err := DetectContentType(data) // Ensure DetectContentType returns the combined reader!
		if err != nil {
			return err
		}
		contentType = &ct
		body = newReader
	}

	// 2. Use the Uploader instead of PutObject
	// The uploader is smart: it buffers parts of the reader into memory
	// so it can calculate checksums for each part individually.
	uploader := manager.NewUploader(s.client)

	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: contentType,
		ACL:         "public-read",
	})

	return err
}

// func (s *S3Client) UploadBytesStream(ctx context.Context, data io.Reader, key string, contentType *string) error {
// 	var body io.Reader = data

// 	// If you are detecting content type, the reader becomes unseekable here
// 	if contentType == nil {
// 		ct, newReader, err := DetectContentType(data)
// 		if err != nil {
// 			return err
// 		}
// 		contentType = &ct
// 		body = newReader
// 	}

// 	// Use the Uploader instead of s.client.PutObject
// 	uploader := manager.NewUploader(s.client)
// 	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
// 		ChecksumAlgorithm: types.ChecksumAlgorithm(""),
// 		Bucket:            aws.String(s.bucket),
// 		Key:               aws.String(key),
// 		Body:              body,
// 		ContentType:       contentType,
// 		ACL:               "public-read",
// 	})

// 	return err
// }

func DetectContentType(r io.Reader) (string, io.Reader, error) {
	buf := make([]byte, 512)
	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return "", nil, err
	}

	ct := http.DetectContentType(buf)
	return ct, io.MultiReader(bytes.NewReader(buf[:n]), r), nil
}

func (s *S3Client) GetPresignedDownloadURL(ctx context.Context, key string, expires *time.Duration) (string, error) {
	req, err := s.presign.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		if expires != nil {
			opts.Expires = *expires
		}
	})

	if err != nil {
		return "", err
	}

	return req.URL, nil
}

func (s *S3Client) GetURL(key string) string {
	videoUrl := fmt.Sprintf("%s/%s/%s", s.config.S3.AWSURL, s.bucket, key)
	return videoUrl
}

func (s *S3Client) GetMetadata(
	ctx context.Context,
	key string,
) (map[string]string, *s3.HeadObjectOutput, error) {

	out, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, nil, err
	}

	// User-defined metadata (x-amz-meta-*)
	// Keys are LOWERCASED by AWS SDK
	return out.Metadata, out, nil
}

// Download file from S3
func (s *S3Client) DownloadFile(ctx context.Context, key string) (io.ReadCloser, error) {
	out, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return out.Body, nil
}

// Delete file from S3
func (s *S3Client) DeleteFile(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}
