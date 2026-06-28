package config

import "os"

type Config struct {
	DatabaseURL string
	JWTSecret   string
	Port        string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://securebank:securebank123@localhost:5432/accountdb?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "local-dev-secret-change-in-prod"),
		Port:        getEnv("PORT", "8081"),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
