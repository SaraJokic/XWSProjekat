package jwt

import (
	"fmt"
	"xwsproj/model"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtClaims struct {
	Id       primitive.ObjectID `bson:"id,omitempty" json:"id"`
	Name     string             `bson:"name"     json:"name"`
	Username string             `bson:"username" json:"username"`
	Role     model.UserType     `bson:"role"     json:"role"`
	jwt.StandardClaims   
}

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) {
		return nil
	}
	return fmt.Errorf("Token is invalid")
}
