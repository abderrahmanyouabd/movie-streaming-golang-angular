package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"movie-streaming-backend/internal/config"
	"movie-streaming-backend/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Load configuration
	cfg := config.LoadConfig()

	// Set Gin mode
	gin.SetMode(cfg.GinMode)

	// MongoDB Connection
	client, err := repository.ConnectDB(cfg.MongoDBURI)
	if err != nil {
		slog.Error("Could not connect to MongoDB", "error", err)
	} else {
		defer func() {
			if err := client.Disconnect(context.Background()); err != nil {
				slog.Error("Failed to disconnect MongoDB", "error", err)
			}
		}()
	}

	// Initialize Gin router
	r := setupRouter()

	slog.Info("Server starting", "port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}

func setupRouter() *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "up",
			"message": "Movie Streaming API with Gin",
		})
	})

	return r
}
