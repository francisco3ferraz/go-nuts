package config

import (
	"os"
	"time"
)

const (
	defaultHTTPAddr         = ":8080"
	defaultEnvironment      = "development"
	defaultShutdownTimeout  = 10 * time.Second
	defaultReadHeaderTimout = 5 * time.Second
)

type Config struct {
	Environment       string
	HTTPAddr          string
	ShutdownTimeout   time.Duration
	ReadHeaderTimeout time.Duration
}

func Load() Config {
	return Config{
		Environment:       getEnv("APP_ENV", defaultEnvironment),
		HTTPAddr:          getEnv("HTTP_ADDR", defaultHTTPAddr),
		ShutdownTimeout:   getDurationEnv("HTTP_SHUTDOWN_TIMEOUT", defaultShutdownTimeout),
		ReadHeaderTimeout: getDurationEnv("HTTP_READ_HEADER_TIMEOUT", defaultReadHeaderTimout),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getDurationEnv(key string, fallback time.Duration) time.Duration {
	raw := os.Getenv(key)
	if raw == "" {
		return fallback
	}

	parsed, err := time.ParseDuration(raw)
	if err != nil {
		return fallback
	}

	return parsed
}
