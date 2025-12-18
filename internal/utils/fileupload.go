package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	s3Client "github.com/sojebsikder/go-boilerplate/pkg/s3client"
)

type FileUploadOpts struct {
	Context  *gin.Context
	FormKey  string
	S3Client s3Client.S3Client
	Path     string
}

func UploadFilesToS3(opt FileUploadOpts) ([]string, error) {
	form, err := opt.Context.MultipartForm()
	if err != nil {
		return nil, fmt.Errorf("failed to parse multipart form: %w", err)
	}
	if form == nil || form.File == nil || len(form.File[opt.FormKey]) == 0 {
		return nil, nil
	}

	var uploadedFiles []string

	for _, fileHeader := range form.File[opt.FormKey] {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %w", fileHeader.Filename, err)
		}
		defer file.Close()

		uploadedFileName := fmt.Sprintf("%s/%s_%s", opt.Path, uuid.New().String(), fileHeader.Filename)

		if err := opt.S3Client.UploadFile(opt.Context.Request.Context(), file, uploadedFileName); err != nil {
			return nil, fmt.Errorf("failed to upload file %s to S3: %w", fileHeader.Filename, err)
		}

		uploadedFiles = append(uploadedFiles, uploadedFileName)
	}

	return uploadedFiles, nil
}
