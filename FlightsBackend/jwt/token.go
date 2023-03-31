package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	jWTPrivateToken = "SecrteTokenSecrteToken"
)

func GenerateToken(claims *JwtClaims, expirationTime time.Time) (string, error) {

	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jWTPrivateToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
