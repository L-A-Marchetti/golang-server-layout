package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

// LoggingMiddleware logs details about incoming requests.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Record start time

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log request details after processing
		duration := time.Since(start)
		fmt.Printf("Method: %s | Path: %s | Duration: %v\n", r.Method, r.URL.Path, duration)
	}
}
