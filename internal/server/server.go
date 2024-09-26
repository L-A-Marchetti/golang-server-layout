package server

import (
	"fmt"
	"handlers"
	"middlewares"
	"time"
)

func InitServer() {
	// Create a new server instance with specified timeout settings and max header bytes
	server := NewServer(":8080", 10*time.Second, 10*time.Second, 30*time.Second, 2*time.Second, 1<<20) // 1 MB max header size

	// Add handlers for different routes
	server.Handle("/", handlers.IndexHandler)      // Root route
	server.Handle("/about", handlers.AboutHandler) // About route

	// Add middlewares
	server.Use(middlewares.LoggingMiddleware)
	server.Use(middlewares.NotFoundMiddleware)

	if err := server.Start(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
