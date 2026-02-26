package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Set Gin to release mode in production, but keep debug for now
	// gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		slog.Info("Request received", "method", c.Request.Method, "path", c.Request.URL.Path)
		c.String(http.StatusOK, "Movie Streaming API with Gin")
	})

	slog.Info("Server listening on :8080")
	if err := r.Run(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
