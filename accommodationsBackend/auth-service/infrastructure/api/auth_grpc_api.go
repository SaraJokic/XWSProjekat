package api

import (
	"accommodationsBackend/auth-service/application"
	"accommodationsBackend/auth-service/domain"
	"accommodationsBackend/auth-service/jwt"
	auth_service "accommodationsBackend/common/proto/auth-service"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AuthHandler struct {
	auth_service.UnimplementedAuthServiceServer
	service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (handler *AuthHandler) Insert(ctx context.Context, request *auth_service.InsertRequest) (*auth_service.InsertResponse, error) {
	newUser := &domain.User{
		//Id:       primitive.NewObjectID(),
		Username: request.Username,
		Password: request.Password,
	}
	fmt.Println("ovo je newuseer u insert handleru", newUser.Username, newUser.Password)
	newUser.Id = primitive.NewObjectID()

	err := handler.service.Insert(newUser)
	if err != nil {
		return nil, err
	}
	response := &auth_service.InsertResponse{Message: "Registration successful!"}
	return response, nil

}

func (handler *AuthHandler) Login(ctx context.Context, request *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {
	/*var loginObj model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginObj)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}
	*/
	var user *domain.User
	user, err := handler.service.ValidateUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	var claims = &jwt.JwtClaims{}
	claims.Id = user.Id
	claims.Name = user.Username
	claims.Username = request.Username

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)
	tokenString, err := jwt.GenerateToken(claims, expirationTime)

	if err != nil {
		return nil, err
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return &auth_service.LoginResponse{Message: string(jsonResponse)}, nil

	/*jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	//response.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)*/
}
func (handler *AuthHandler) GetAll(ctx context.Context, request *auth_service.AllRequest) (*auth_service.AllResponse, error) {
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &auth_service.AllResponse{
		Users: []*auth_service.AuthUser{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
