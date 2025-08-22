package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	SQLitePath string
	JWTSecret  string
}

// Load reads environment variables (loading .env if present) and returns Config with defaults.
func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		Port:       getEnv("PORT", "8080"),
		SQLitePath: getEnv("SQLITE_PATH", "./app.db"),
		JWTSecret:  getEnv("JWT_SECRET", "devsecret"),
	}
	return cfg
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
