package middleware

import (
	"context"
	"net/http"
	"strings"
	"xwsproj/jwt"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Empty string", http.StatusUnauthorized)
			return
		}
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {      
			http.Error(w, "split didnt work", http.StatusUnauthorized)
			return
		}
		tokenString := headerParts[1]

		valid, claims := jwt.VerifyToken(tokenString)
		if !valid {
			http.Error(w, "token not verified", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "Id", claims.Id)
		ctx = context.WithValue(ctx, "Name", claims.Name)
		ctx = context.WithValue(ctx, "Username", claims.Username)
		ctx = context.WithValue(ctx, "Role", claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
