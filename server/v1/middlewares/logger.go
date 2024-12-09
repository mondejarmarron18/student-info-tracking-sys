package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom response writer to capture the status code
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Log the HTTP method, URL path, and status code
		log.Println(r.Method, r.URL.Path, colorizeStatusCode(rw.statusCode))
	})
}

func colorizeStatusCode(code int) string {
	switch {
	case code >= 200 && code < 300:
		return fmt.Sprintf("\033[32m%d\033[0m", code) // Green for 2xx
	case code >= 300 && code < 400:
		return fmt.Sprintf("\033[36m%d\033[0m", code) // Cyan for 3xx
	case code >= 400 && code < 500:
		return fmt.Sprintf("\033[33m%d\033[0m", code) // Yellow for 4xx
	case code >= 500:
		return fmt.Sprintf("\033[31m%d\033[0m", code) // Red for 5xx
	default:
		return fmt.Sprintf("%d", code) // Default color
	}
}
