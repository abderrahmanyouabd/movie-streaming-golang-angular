package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"movie-streaming-backend/internal/config"
	"movie-streaming-backend/internal/handler"
	"movie-streaming-backend/internal/repository"
	"movie-streaming-backend/internal/service"

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

	// Database instance
	db := client.Database(cfg.DBName)

	// Dependency Injection
	movieRepo := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepo)
	movieHandler := handler.NewMovieHandler(movieService)

	// Initialize Gin router
	r := setupRouter(movieHandler)

	slog.Info("Server starting", "port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}

func setupRouter(movieHandler *handler.MovieHandler) *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "up",
			"message": "Movie Streaming API with Gin",
		})
	})

	// Movies Group
	api := r.Group("/api/v1")
	{
		movies := api.Group("/movies")
		{
			movies.POST("", movieHandler.AddMovie)
			movies.GET("", movieHandler.GetAllMovies)
		}
	}

	return r
}
