package config

import (
	"os"
	"strings"
)

type Config struct {
	Environment    string
	Port          string
	AllowedOrigins []string
	DatabaseURL   string
}

func Load() *Config {
	return &Config{
		Environment:    getEnv("ENV", "development"),
		Port:          getEnv("PORT", "8080"),
		AllowedOrigins: strings.Split(getEnv("ALLOWED_ORIGINS", "http://localhost:3000"), ","),
		DatabaseURL:   getEnv("DATABASE_URL", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}