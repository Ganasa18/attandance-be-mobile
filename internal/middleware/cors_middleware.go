package middleware

import (
	"net/http"
	"time"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewBaseMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Set CORS headers
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	writer.Header().Set("Cache-Control", "no-cache, max-age=120, must-revalidate")

	// Set Expires header to a future time to indicate the expiration time
	expiresTime := time.Now().Add(120 * time.Second)
	writer.Header().Set("Expires", expiresTime.UTC().Format(http.TimeFormat))

	// Set Last-Modified header to the current time
	writer.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))

	// Handle preflight requests
	if request.Method == http.MethodOptions {
		writer.WriteHeader(http.StatusOK)
		return
	}

	// Call the next handler
	middleware.Handler.ServeHTTP(writer, request)
}
