package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/*
const (

	jWTTokenKey = "SecretTokenKey"

)

func GenerateToken(claims *JwtAccessTokenClaims, expirationTime time.Time) (string, error) {

		claims.ExpiresAt = expirationTime.Unix()
		claims.IssuedAt = time.Now().UTC().Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(jWTTokenKey))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
*/
const (
	tokenSecretKey = "SecrteTokenSecrteToken"
)

// Kreiranje JWT tokena na osnovu pruženih tvrdnji (claimData) i vremena isteka
func GenerateToken(claimData *JwtClaims) (string, error) {
	currentTime := time.Now().UTC()
	claimData.IssuedAt = currentTime.Unix()
	claimData.ExpiresAt = currentTime.Add(2 * time.Hour).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimData)

	tokenString, err := jwtToken.SignedString([]byte(tokenSecretKey))
	if err != nil {

		return "", fmt.Errorf("failed to sign the JWT with the secret key: %v", err)
	}

	return tokenString, nil
}
