package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v11"
)

// Config holds application configuration parsed from environment variables.
// Supports Docker Swarm-style secrets via *_FILE env vars (auto-detected).
// Example: Set DB_PASSWORD_FILE=/run/secrets/db_pass to load secret from file.
type Config struct {
	// App settings
	AppName         string `env:"APP_NAME" envDefault:"Fiber Boilerplate"`
	Port            string `env:"PORT" envDefault:"3000"`
	ShutdownTimeout int    `env:"SHUTDOWN_TIMEOUT" envDefault:"5"`

	// Secrets (supports _FILE override, e.g., APP_SECRET_FILE=/path/to/secret)
	AppSecret string `env:"APP_SECRET"`

	// Database (example for future use)
	// DatabaseURL string `env:"DATABASE_URL,required"`
	// DBPassword  string `env:"DB_PASSWORD"`
}

// preloadFileSecrets scans all environment variables for *_FILE suffix.
// If found, reads the file contents and sets the base variable.
// Example: DB_PASSWORD_FILE=/run/secrets/db_pass → sets DB_PASSWORD=<file contents>
func preloadFileSecrets() error {
	const suffix = "_FILE"
	for _, environ := range os.Environ() {
		pair := strings.SplitN(environ, "=", 2)
		if len(pair) != 2 {
			continue
		}
		key, value := pair[0], pair[1]

		if !strings.HasSuffix(key, suffix) || value == "" {
			continue
		}

		baseKey := strings.TrimSuffix(key, suffix)
		content, err := os.ReadFile(value)
		if err != nil {
			return fmt.Errorf("failed to read %s from %s: %w", key, value, err)
		}

		os.Setenv(baseKey, strings.TrimSpace(string(content)))
	}
	return nil
}

// InitConfig parses environment variables into Config struct.
// Automatically loads *_FILE secrets before parsing.
// Returns an error if required fields are missing or file reading fails.
func InitConfig() (*Config, error) {
	// Auto-load all *_FILE secrets first
	if err := preloadFileSecrets(); err != nil {
		return nil, err
	}

	// Parse env vars into struct
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}

// Must wraps a value and error, panicking if error is non-nil.
// Useful for config initialization in main() where failure should halt startup.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("config error: %v", err))
	}
	return val
}
