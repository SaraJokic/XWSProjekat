package api

import (
	"accommodationsBackend/accommodations-service/application"
	"accommodationsBackend/common/proto/accommodation_service"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationHandler struct {
	accommodation_service.UnimplementedAccommodationServiceServer
	service *application.AccommodationService
}

func NewAccommodationHandler(service *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (handler *AccommodationHandler) Get(ctx context.Context, request *accommodation_service.AccGetRequest) (*accommodation_service.AccGetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	accommodation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	accMapped := mapAccommodation(accommodation)
	response := &accommodation_service.AccGetResponse{
		Acc: accMapped,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *accommodation_service.AccGetAllRequest) (*accommodation_service.AccGetAllResponse, error) {
	accommodations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &accommodation_service.AccGetAllResponse{
		Acc: []*accommodation_service.Accommodation{},
	}
	for _, accommodation := range accommodations {
		current := mapAccommodation(accommodation)
		response.Acc = append(response.Acc, current)
	}
	return response, nil
}
func (handler *AccommodationHandler) CreateNewAccommodation(ctx context.Context, request *accommodation_service.AccCreateRequest) (*accommodation_service.AccCreateResponse, error) {

	acc := mapNewAccommodation(request.Acc)

	err := handler.service.Create(acc)
	if err != nil {
		return nil, err
	}

	accommodation := accommodation_service.AccCreateResponse{Id: acc.Id.Hex()}

	return &accommodation, nil
}
