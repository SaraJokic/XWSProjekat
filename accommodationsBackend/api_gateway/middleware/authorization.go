package middleware

import (
	"accommodationsBackend/api_gateway/jwt"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

/*
func AuthorizeAndAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/auth/login" || r.URL.Path == "/auth/insert" || r.URL.Path == "/users/register" || r.URL.Path == "/accommodations" {
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
		ctx = context.WithValue(ctx, "Role", claims.Role)

		fmt.Println(claims.Role)

		var isAuthorized bool = true
		switch r.URL.Path {
		case "/users":
			isAuthorized = Authorization("Customer")(w, r.WithContext(ctx))
		case "/GetAllProminentHosts":
			isAuthorized = Authorization("Admin")(w, r.WithContext(ctx))
		default:
			isAuthorized = true
		}

		if !isAuthorized {
			return
		}

		//md := metadata.Pairs("authorization", "Bearer "+tokenString)
		//ctx = metadata.NewOutgoingContext(r.Context(), md)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Authorization(validRole string) func(http.ResponseWriter, *http.Request) bool {
	return func(w http.ResponseWriter, r *http.Request) bool {
		fmt.Println("usao sam u metodu")
		roleVal := r.Context().Value("Role")
		fmt.Println("ovo sam procitao kao rolu", roleVal)

		if roleVal == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return false
		}

		role, ok := roleVal.(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return false
		}

		fmt.Println("role", role)

		if role != validRole {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return false
		}
		return true
	}
}
*/

func AuthorizeAndAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAuthorized := true
		if r.URL.Path == "/auth/login" || r.URL.Path == "/auth/insert" || r.URL.Path == "/users/register" || r.URL.Path == "/accommodations" || r.URL.Path == "/accommodations/search" {
			next.ServeHTTP(w, r)
			return
		}
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)

			return
		}

		tokenString := getTokenStringFromHeader(authorizationHeader)
		if tokenString == "" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		valid, claims := jwt.VerifyToken(tokenString)
		if !valid {
			http.Error(w, "ovo je posle metode verify token Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := addClaimsToContext(r.Context(), claims)
		isAuthorized = checkAuthorization(r.URL.Path, w, r.WithContext(ctx))
		if !isAuthorized {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getTokenStringFromHeader(header string) string {
	parts := strings.Split(header, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}

func addClaimsToContext(ctx context.Context, claims *jwt.JwtClaims) context.Context {

	ctx = context.WithValue(ctx, "Id", claims.Id)
	ctx = context.WithValue(ctx, "Name", claims.Name)
	ctx = context.WithValue(ctx, "Username", claims.Username)
	return context.WithValue(ctx, "Role", claims.Role)
}

func checkAuthorization(path string, w http.ResponseWriter, r *http.Request) bool {
	switch path {
	case "/reservations/update", "/reservations/create", "/reservations/delete/{id}", "/ratings/addhost", "/ratings/addacc":
		return Authorization("Customer")(w, r)

	case "/accommodations/create", "/accommodations/update", "/availability/addpricechange", "/availability/addavailableslot", "/availability/update/availableslot", "/availability/add",
		"/availability/update", "/availability/updateafter", "/reservations/changestatus":
		return Authorization("Admin")(w, r)

	case "/users/delete/{id}", "/users/update":
		return Authorization("Customer")(w, r) || Authorization("Admin")(w, r)

	default:
		return true
	}
}

func Authorization(validRole string) func(http.ResponseWriter, *http.Request) bool {
	return func(w http.ResponseWriter, r *http.Request) bool {
		roleVal := r.Context().Value("Role")

		role, ok := roleVal.(string)
		if !ok || roleVal == nil || role != validRole {
			return false
		}
		return true
	}
}
