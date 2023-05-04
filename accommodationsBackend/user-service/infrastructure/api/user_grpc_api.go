package api

import (
	"accommodationsBackend/common/proto/user_service"
	"accommodationsBackend/user-service/application"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	user_service.UnimplementedUserServiceServer
	service *application.UserService
}

func NewProductHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *user_service.GetRequest) (*user_service.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	userMapped := mapUser(user)
	response := &user_service.GetResponse{
		User: userMapped,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *user_service.GetAllRequest) (*user_service.GetAllResponse, error) {
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &user_service.GetAllResponse{
		Users: []*user_service.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
