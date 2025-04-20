package middleware

import (
	"backend/config"
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Define custom types for context keys to avoid collisions
type contextKey string

const (
	contextKeyAdminID contextKey = "admin_id"
	contextKeyRole    contextKey = "role"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	log.Printf("JWT Middleware initialized")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] Request: %s %s", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)

		// Bypass GET requests and authentication routes
		if r.Method == http.MethodGet || strings.HasPrefix(r.URL.Path, "/api/admins/authenticate") {
			next.ServeHTTP(w, r)
			return
		}

		// Extract the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Get the JWT secret from environment variables
		secret := config.GetEnv("JWT_SECRET", "")
		if secret == "" {
			http.Error(w, "JWT secret not configured", http.StatusInternalServerError)
			return
		}

		// Parse and validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract claims and add them to the request context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		adminID, ok := claims["admin_id"].(float64)
		if !ok {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add admin ID and role to the request context
		r = r.WithContext(context.WithValue(r.Context(), contextKeyAdminID, uint(adminID)))
		r = r.WithContext(context.WithValue(r.Context(), contextKeyRole, role))

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}