package s3client

import (
	"go.uber.org/fx"
)

var Module = fx.Module("modules", fx.Provide(
	NewS3Client,
))
