package jwt

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

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
