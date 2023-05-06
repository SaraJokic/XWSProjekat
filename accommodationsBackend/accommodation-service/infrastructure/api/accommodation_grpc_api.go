package api

import (
	"accommodationsBackend/accommodations-service/application"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationHandler struct {
	accommodation_service.UnimplementedAccommodationServiceServer
	service *application.AccommodationService
}

func NewProductHandler(service *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (handler *AccommodationHandler) Get(ctx context.Context, request *accommodation_service.GetRequest) (*accommodation_service.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	userMapped := accMapped(user)
	response := &accommodation_service.GetResponse{
		User: userMapped,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *accommodation_service.GetAllRequest) (*accommodation_service.GetAllResponse, error) {
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &uaccommodation_service.GetAllResponse{
		Users: []*accommodation_service.User{},
	}
	for _, user := range users {
		current := accMapped(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
