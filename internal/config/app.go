package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type Config struct {
	App       AppConfig
	Database  DatabaseConfig
	S3        AWSConfig
	Redis     RedisConfig
	Security  SecurityConfig
	Mail      MailConfig
	RateLimit RateLimitConfig
}

type AppConfig struct {
	Name string `env:"APP_NAME,default=go-boilerplate"`
	Port string `env:"PORT,default=8080"`
	// TemplateDir stores the name of the directory that contains templates
	TemplateDir string `env:"TEMPLATE_DIR,default=views"`
	// StaticDir stores the name of the directory that will serve static files
	StaticDir string `env:"STATIC_DIR,default=static"`
}

type DatabaseConfig struct {
	DatabaseURL string `env:"DATABASE_URL,default=postgres://user:pass@localhost:5432/testdemo"`
}

type AWSConfig struct {
	AWSAccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	AWSRegion          string `env:"AWS_REGION,default=us-east-1"`
	AWSBucket          string `env:"AWS_BUCKET"`
	AWSURL             string `env:"AWS_URL"`
	AWSEndpoint        string `env:"AWS_ENDPOINT"`
}

type RedisConfig struct {
	RedisURL string `env:"REDIS_URL,default=localhost:6379"`
	Password string `env:"REDIS_PASSWORD,default=secret"`
}

type SecurityConfig struct {
	JWTSecret string `env:"JWT_SECRET,default=secret"`
}

type MailConfig struct {
	Hostname    string `env:"MAIL_HOST,default=localhost"`
	Port        uint16 `env:"MAIL_PORT,default=25"`
	User        string `env:"MAIL_USERNAME,default=admin"`
	Password    string `env:"MAIL_PASSWORD,default=admin"`
	FromAddress string `env:"MAIL_FROM_ADDRESS,default=admin@localhost"`
}

type RateLimitConfig struct {
	RateLimitMaxRequests int64         `env:"RATE_LIMIT_MAX_REQUESTS,default=10"`
	RateLimitDuration    time.Duration `env:"RATE_LIMIT_DURATION,default=1m"`
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config
	err = envdecode.StrictDecode(&cfg)
	return &cfg, err
}

var Module = fx.Module("config", fx.Provide(NewConfig))
