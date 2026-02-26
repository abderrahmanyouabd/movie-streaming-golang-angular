package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	MongoDBURI string
	DBName     string
	GinMode    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("No .env file found, using system environment variables")
	}

	return &Config{
		Port:       getEnv("PORT", "8080"),
		MongoDBURI: getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		DBName:     getEnv("DB_NAME", "movie_streaming"),
		GinMode:    getEnv("GIN_MODE", "debug"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
