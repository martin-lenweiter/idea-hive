package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func CorsMiddleware() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://ideahive.io"}, // Production domain
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			// Allow localhost in development
			return origin == "http://localhost:8080" || origin == "http://127.0.0.1:8080"
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value for Access-Control-Max-Age
	})
}
