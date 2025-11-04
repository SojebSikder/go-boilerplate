package config

import (
	"log"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Security SecurityConfig
	Mail     MailConfig
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

type SecurityConfig struct {
	JWTKey string `env:"JWT_KEY,default=secret"`
}

type MailConfig struct {
	Hostname    string `env:"MAIL_HOST,default=localhost"`
	Port        uint16 `env:"MAIL_PORT,default=25"`
	User        string `env:"MAIL_USERNAME,default=admin"`
	Password    string `env:"MAIL_PASSWORD,default=admin"`
	FromAddress string `env:"MAIL_FROM_ADDRESS,default=admin@localhost"`
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
