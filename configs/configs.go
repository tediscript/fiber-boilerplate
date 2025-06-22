package configs

import (
	"os"
	"strconv"
)

type Config struct {
	AppName         string
	Port            string
	ShutdownTimeout int
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}

func InitConfig() *Config {
	return &Config{
		AppName:         getEnv("APP_NAME", "Fiber Boilerplate"),
		Port:            getEnv("PORT", "3000"),
		ShutdownTimeout: getEnvAsInt("SHUTDOWN_TIMEOUT", 5),
	}
}
