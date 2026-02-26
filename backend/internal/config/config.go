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
	// Check for .env in current dir or parent (root)
	err := godotenv.Load(".env", "../.env")
	if err != nil {
		slog.Warn("No .env file found, using system environment variables")
	}

	user := getEnv("MONGO_USER", "root")
	pass := getEnv("MONGO_PASS", "rootpassword")
	port := getEnv("MONGO_PORT", "27017")

	// Fallback to MONGODB_URI from .env if present and explicitly set (without literal ${ variables)
	mongoURI := getEnv("MONGODB_URI", "")
	if mongoURI == "" || mongoURI == "mongodb://${MONGO_USER}:${MONGO_PASS}@localhost:${MONGO_PORT}/?authSource=admin" {
		mongoURI = "mongodb://" + user + ":" + pass + "@localhost:" + port + "/?authSource=admin"
	}

	return &Config{
		Port:       getEnv("PORT", "8080"),
		MongoDBURI: mongoURI,
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
