package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Initialize structured logger (JSON for production, Text for dev)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request received", "method", r.Method, "path", r.URL.Path)
		w.Write([]byte("Movie Streaming API"))
	})

	slog.Info("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
