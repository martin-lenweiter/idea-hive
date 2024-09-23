package middleware

import (
	"log"
	"net/http"
	"os"

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

func ForceSSL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("ENV", os.Getenv("ENV"))
		log.Println("X-Forwarded-Proto", r.Header.Get("X-Forwarded-Proto"))
		if os.Getenv("ENV") == "production" && r.Header.Get("X-Forwarded-Proto") != "https" {
			// Redirect to HTTPS in production
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r) // Continue for local development
	})
}
