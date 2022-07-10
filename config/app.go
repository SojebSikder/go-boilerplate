package config

import (
	"os"

	"github.com/joeshaw/envdecode"
)

const (
	// TemplateDir stores the name of the directory that contains templates
	TemplateDir = "views"

	// TemplateExt stores the extension used for the template files
	TemplateExt = ".gohtml"

	// StaticDir stores the name of the directory that will serve static files
	StaticDir = "static"

	// StaticPrefix stores the URL prefix used when serving static files
	StaticPrefix = "files"
)

type environment string

const (
	// EnvLocal represents the local environment
	EnvLocal environment = "local"

	// EnvTest represents the test environment
	EnvTest environment = "test"

	// EnvDevelop represents the development environment
	EnvDevelop environment = "dev"

	// EnvStaging represents the staging environment
	EnvStaging environment = "staging"

	// EnvQA represents the qa environment
	EnvQA environment = "qa"

	// EnvProduction represents the production environment
	EnvProduction environment = "prod"
)

// SwitchEnvironment sets the environment variable used to dictate which environment the application is
// currently running in.
// This must be called prior to loading the configuration in order for it to take effect.
func SwitchEnvironment(env environment) {
	if err := os.Setenv("APP_ENVIRONMENT", string(env)); err != nil {
		panic(err)
	}
}

type (
	// Config stores complete configuration
	Config struct {
		Database DatabaseConfig
	}
	// DatabaseConfig stores the database configuration
	DatabaseConfig struct {
		Hostname     string `env:"DB_HOST,default=localhost"`
		Port         uint16 `env:"DB_PORT,default=5432"`
		User         string `env:"DB_USERNAME,default=admin"`
		Password     string `env:"DB_PASSWORD,default="`
		Database     string `env:"DB_DATABASE,default=app"`
		TestDatabase string `env:"DB_DATABASE_TEST,default=app_test"`
	}
)

// GetConfig loads and returns configuration
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg)
	return cfg, err
}
