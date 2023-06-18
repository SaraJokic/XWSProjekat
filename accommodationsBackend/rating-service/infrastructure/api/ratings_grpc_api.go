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
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	hosts, err := handler.service.GetHostRatingsByGuestId(objectId)
	if err != nil {
		return nil, err
	}
	response := &rating_service.GetRateHostByGuestResponse{
		Response: []*rating_service.RateHost{},
	}
	for _, a := range hosts {
		current := mapHostRating(a)
		response.Response = append(response.Response, current)
	}
	return response, nil
}
func (handler *RatingsHandler) GetAccommodationsRatingsByGuestId(ctx context.Context, request *rating_service.GetRateAccommodationByGuestRequest) (*rating_service.GetRateAccommodationByGuestResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	accommodations, err := handler.service.GetAccommodationsRatingsByGuestId(objectId)
	if err != nil {
		return nil, err
	}
	response := &rating_service.GetRateAccommodationByGuestResponse{
		Response: []*rating_service.RateAccommodation{},
	}
	for _, a := range accommodations {
		current := mapAccommodationRating(a)
		response.Response = append(response.Response, current)
	}
	return response, nil
}
func (handler *RatingsHandler) GetAccommodationsRatingsByAccommodationId(ctx context.Context, request *rating_service.GetRateAccommodationByAccommodationRequest) (*rating_service.GetRateAccommodationByAccommodationResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	accommodations, err := handler.service.GetAccommodationsRatingsByAccommodationId(objectId)
	if err != nil {
		return nil, err
	}
	response := &rating_service.GetRateAccommodationByAccommodationResponse{
		Response: []*rating_service.RateAccommodation{},
	}
	for _, a := range accommodations {
		current := mapAccommodationRating(a)
		response.Response = append(response.Response, current)
	}
	return response, nil
}
func (handler *RatingsHandler) GetHostRatingsByHostId(ctx context.Context, request *rating_service.GetRateHostByHostRequest) (*rating_service.GetRateHostByHostResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	hosts, err := handler.service.GetHostRatingsByHostId(objectId)
	if err != nil {
		return nil, err
	}
	response := &rating_service.GetRateHostByHostResponse{
		Response: []*rating_service.RateHost{},
	}
	for _, a := range hosts {
		current := mapHostRating(a)
		response.Response = append(response.Response, current)
	}
	return response, nil
}
func (handler *RatingsHandler) UpdateAccommodationRating(ctx context.Context, request *rating_service.UpdateAccommodationRatingRequest) (*rating_service.UpdateAccommodationRatingResponse, error) {
	id := request.Id
	accommodation := request.Rating
	accMapped := mapUpdateAccommodationRating(accommodation)
	err := handler.service.UpdateAccommodationRating(id, accMapped)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update accommodation")
	}

	return &rating_service.UpdateAccommodationRatingResponse{Response: accommodation}, nil
}
func (handler *RatingsHandler) UpdateHostRating(ctx context.Context, request *rating_service.UpdateHostRatingRequest) (*rating_service.UpdateHostRatingResponse, error) {
	id := request.Id
	host := request.Rating
	hostMapped := mapUpdateHostRating(host)
	err := handler.service.UpdateHostRating(id, hostMapped)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update accommodation")
	}

	return &rating_service.UpdateHostRatingResponse{Response: host}, nil
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
func (handler *RatingsHandler) GetAvgRatingHost(ctx context.Context, request *rating_service.GetAvgHostRatingRequest) (*rating_service.GetAvgHostRatingResponse, error) {
	hosts, _ := handler.GetHostRatingsByHostId(ctx, &rating_service.GetRateHostByHostRequest{Id: request.Id})
	count := 0
	sum := 0.0
	for _, rating := range hosts.Response {
		sum += float64(rating.Rating)
		count++
	}
	if count > 0 {
		return &rating_service.GetAvgHostRatingResponse{
			Avg: float32(sum / float64(count)),
		}, nil
	}
	return &rating_service.GetAvgHostRatingResponse{
		Avg: 0.0,
	}, nil

}
func (handler *RatingsHandler) GetAvgAccommodationRating(ctx context.Context, request *rating_service.GetAvgAccommodationRatingRequest) (*rating_service.GetAvgAccommodationRatingResponse, error) {
	accommodations, _ := handler.GetAccommodationsRatingsByAccommodationId(ctx, &rating_service.GetRateAccommodationByAccommodationRequest{Id: request.Id})
	count := 0
	sum := 0.0
	for _, rating := range accommodations.Response {
		sum += float64(rating.Rating)
		count++
	}
	if count > 0 {
		return &rating_service.GetAvgAccommodationRatingResponse{
			Avg: float32(sum / float64(count)),
		}, nil
	}
	return &rating_service.GetAvgAccommodationRatingResponse{
		Avg: 0.0,
	}, nil
}
