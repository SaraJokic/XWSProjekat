package middleware

import (
	"accommodationsBackend/api_gateway/jwt"
	"context"
	"net/http"
	"strings"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/auth/login" || r.URL.Path == "/auth/insert" || r.URL.Path == "/users/register" {
			next.ServeHTTP(w, r)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)

			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)

			return
		}

		tokenString := headerParts[1]

		valid, claims := jwt.VerifyToken(tokenString)
		if !valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)

			return
		}

		ctx := context.WithValue(r.Context(), "Id", claims.Id)
		ctx = context.WithValue(ctx, "Name", claims.Name)
		ctx = context.WithValue(ctx, "Username", claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
