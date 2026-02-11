package s3client

import (
	"go.uber.org/fx"
)

var Module = fx.Module("s3client", fx.Provide(
	NewS3Client,
))
