package api

import (
	"accommodationsBackend/auth-service/application"
	"accommodationsBackend/auth-service/domain"
	auth_service "accommodationsBackend/common/proto/auth-service"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
