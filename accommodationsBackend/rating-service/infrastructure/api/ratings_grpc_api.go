package api

import (
	"accommodationsBackend/common/proto/rating_service"
	"accommodationsBackend/rating-service/application"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RatingsHandler struct {
	rating_service.UnimplementedRatingServiceServer
	service *application.RatingsService
}

func NewRatingsHandler(service *application.RatingsService) *RatingsHandler {
	return &RatingsHandler{
		service: service,
	}
}

func (handler *RatingsHandler) GetRateAccommodation(ctx context.Context, request *rating_service.GetRateAccommodationRequest) (*rating_service.GetRateAccommodationResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	rating, err := handler.service.GetRateAccommodation(objectId)
	if err != nil {
		return nil, err
	}
	ratingMapped := mapAccommodationRating(rating)
	response := &rating_service.GetRateAccommodationResponse{
		Response: ratingMapped,
	}
	return response, nil
}
func (handler *RatingsHandler) GetRateHost(ctx context.Context, request *rating_service.GetRateHostRequest) (*rating_service.GetRateHostResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	rating, err := handler.service.GetRateHost(objectId)
	if err != nil {
		return nil, err
	}
	ratingMapped := mapHostRating(rating)
	response := &rating_service.GetRateHostResponse{
		Response: ratingMapped,
	}
	return response, nil
}
func (handler *RatingsHandler) GetHostRatingsByGuestId(ctx context.Context, request *rating_service.GetRateHostByGuestRequest) (*rating_service.GetRateHostByGuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostRatingsByGuestId not implemented")
}
func (handler *RatingsHandler) GetAccommodationsRatingsByGuestId(ctx context.Context, request *rating_service.GetRateAccommodationByGuestRequest) (*rating_service.GetRateAccommodationByGuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationsRatingsByGuestId not implemented")
}
func (handler *RatingsHandler) GetAccommodationsRatingsByAccommodationId(ctx context.Context, request *rating_service.GetRateAccommodationByAccommodationRequest) (*rating_service.GetRateAccommodationByAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationsRatingsByAccommodationId not implemented")
}
func (handler *RatingsHandler) GetHostRatingsByHostId(ctx context.Context, request *rating_service.GetRateHostByHostRequest) (*rating_service.GetRateHostByHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostRatingsByHostId not implemented")
}
func (handler *RatingsHandler) UpdateAccommodationRating(ctx context.Context, request *rating_service.UpdateAccommodationRatingRequest) (*rating_service.UpdateAccommodationRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccommodationRating not implemented")
}
func (handler *RatingsHandler) UpdateHostRating(ctx context.Context, request *rating_service.UpdateHostRatingRequest) (*rating_service.UpdateHostRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHostRating not implemented")
}
func (handler *RatingsHandler) CreateNewHostRating(ctx context.Context, request *rating_service.CreateNewHostRatingRequest) (*rating_service.CreateNewHostRatingResponse, error) {
	acc := mapNewHostRating(request)
	err := handler.service.CreateRateHost(acc)
	if err != nil {
		return nil, err
	}

	rateHost := rating_service.CreateNewHostRatingResponse{Rating: mapHostRating(acc)}

	return &rateHost, nil
}
func (handler *RatingsHandler) CreateNewAccommodationRating(ctx context.Context, request *rating_service.CreateNewAccommodationRatingRequest) (*rating_service.CreateNewAccommodationRatingResponse, error) {
	acc := mapNewAccommodationRating(request)
	err := handler.service.CreateRateAccommodation(acc)
	if err != nil {
		return nil, err
	}

	rateAcc := rating_service.CreateNewAccommodationRatingResponse{Rating: mapAccommodationRating(acc)}

	return &rateAcc, nil
}
