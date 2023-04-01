package jwt

import (
	"context"
	"net/http"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("apikey")

		valid, claims := VerifyToken(tokenString)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "Id", claims.Id)
		ctx = context.WithValue(ctx, "Username", claims.Username)
		ctx = context.WithValue(ctx, "Name", claims.Name)
		ctx = context.WithValue(ctx, "Role", claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
