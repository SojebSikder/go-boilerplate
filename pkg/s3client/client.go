package s3client

import (
	"context"
	"io"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	cfg "github.com/sojebsikder/go-boilerplate/internal/config"
)

type S3Client struct {
	client  *s3.Client
	presign *s3.PresignClient
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
