package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"xwsproj/jwt"
	"xwsproj/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginObj model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate the loginObj for valid credential adn if these are valid then

	var claims = &jwt.JwtClaims{}
	//claims.Id = "ComapnyId"
	claims.Id, _ = primitive.ObjectIDFromHex("ComapnyId")
	claims.Username = loginObj.Username
	claims.Role = 0
	//claims.Audience = context.Request.Header.Get("Referer") // get it from Referer header

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)
	tokenString, err := jwt.GenerateToken(claims, expirationTime)

	if err != nil {
		http.Error(w, "Error in generating token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Token created: %v", tokenString)
}
