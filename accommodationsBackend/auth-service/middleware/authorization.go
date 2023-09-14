package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

const (
	tokenSecretKey = "SecrteTokenSecrteToken" // privatni kljuc
)

/*
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

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
*/
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Ako je zahtev za specifičnu putanju, proverava se token
	if info.FullMethod == "/AuthService/Insert" {
		// Преузимање токена из заглавља
		token, err := extractTokenFromContext(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "no authorization metadata: %v", err)
		}

		// Верификација токена и провера issuer-а
		fmt.Println("metoda koja je", info.FullMethod)
		fmt.Println("usao sam u proveru ZA INSERT U INTERCEPTORY", token)
		isValid, _ := VerifyToken(token, "user-service")
		if !isValid {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token or issuer")
		}
	}

	// Наставите са редовном обрадом
	return handler(ctx, req)
}

func extractTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("no metadata in context")
	}

	authorizationHeaders, ok := md["authorization"]
	if !ok || len(authorizationHeaders) != 1 {
		return "", errors.New("no authorization header")
	}

	tokenParts := strings.Split(authorizationHeaders[0], " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return tokenParts[1], nil
}
func VerifyToken(tokenString, expectedIssuer string) (bool, *JwtClaims) {
	claims := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Header["alg"] != "HS256" {
			return nil, fmt.Errorf("Unexpected signing method or algorithm: %v", token.Header["alg"])
		}
		return []byte(tokenSecretKey), nil
	})

	if err != nil || !token.Valid || claims.Valid() != nil || claims.Issuer != expectedIssuer {
		return false, claims
	}

	return true, claims
}

type JwtClaims struct {
	Id       primitive.ObjectID `bson:"id,omitempty" json:"id"`
	Name     string             `bson:"name"     json:"name"`
	Username string             `bson:"username" json:"username"`
	Role     string             `bson:"role" json:"role"`

	jwt.StandardClaims
}

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) {
		return nil
	}
	return fmt.Errorf("Token is invalid")
}
