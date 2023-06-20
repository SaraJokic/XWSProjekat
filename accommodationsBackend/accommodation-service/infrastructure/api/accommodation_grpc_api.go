package api

import (
	"accommodationsBackend/accommodations-service/application"
	"accommodationsBackend/common/proto/accommodation_service"
	availability_service "accommodationsBackend/common/proto/availability-service"
	"accommodationsBackend/common/proto/user_service"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"
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

func (handler *AccommodationHandler) GetAllProminentAccommodation(ctx context.Context, request *accommodation_service.AccGetAllRequest) (*accommodation_service.AccGetAllResponse, error) {
	accommodations, err := handler.service.GetAllProminentAccommodation()
	if err != nil {
		return nil, err
	}

	response := &accommodation_service.AccGetAllResponse{
		Acc: []*accommodation_service.Accommodation{},
	}
	client := NewUserClient()
	for _, accommodation := range accommodations {
		host, err := client.Get(ctx, &user_service.GetRequest{Id: accommodation.HostId.Hex()})
		if err != nil {

			continue
		}

		if host.User.ProminentHost {
			current := mapAccommodation(accommodation)
			response.Acc = append(response.Acc, current)
		}
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
func (handler *AccommodationHandler) GetByUserId(ctx context.Context, request *accommodation_service.AccGetByUserIdRequest) (*accommodation_service.AccGetByUserIdResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	accommodations, err := handler.service.GetAccommodationByUserId(objectId)
	if err != nil {
		return nil, err
	}
	response := &accommodation_service.AccGetByUserIdResponse{
		Acc: []*accommodation_service.Accommodation{},
	}
	for _, a := range accommodations {
		current := mapAccommodation(a)
		response.Acc = append(response.Acc, current)
	}
	return response, nil
}
func (handler *AccommodationHandler) UpdateAccommodation(ctx context.Context, request *accommodation_service.UpdateAccommodationRequest) (*accommodation_service.UpdateAccommodationResponse, error) {
	id := request.Id
	accommodation := request.Accommodation
	accMapped := mapUpdateAccommodation(accommodation)
	err := handler.service.UpdateAccommodation(id, accMapped)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update accommodation")
	}

	return &accommodation_service.UpdateAccommodationResponse{Accommodation: accommodation}, nil
}
func NewAvailabilityClient() availability_service.AvailabilityServiceClient {
	conn, err := grpc.Dial("availability-service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return availability_service.NewAvailabilityServiceClient(conn)
}
func (handler *AccommodationHandler) Search(ctx context.Context, request *accommodation_service.SearchRequest) (*accommodation_service.AccGetAllResponse, error) {
	location := request.Location
	//guests := request.Guests
	//startDate := request.StartDate
	//endDate := request.EndDate

	//startDateParsed, err := parseISO8601(startDate)
	//endDateParsed, err := parseISO8601(endDate)

	//accommodationsDate, err := handler.service.GetAccommodationByDateRange(startDateParsed, endDateParsed)
	//	if err != nil {
	//return nil, status.Error(codes.Internal, "Failed to update accommodation")

	//}

	accommodationsLocation, err := handler.service.GetAccommodationByLocation(location)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update accommodation")

	}
	/*
		accommodationsGuests, err := handler.service.GetAccommodationByNumberOfGuests(int(guests))
		if err != nil {
			return nil, status.Error(codes.Internal, "Failed to update accommodation")

		}
	*/

	if err != nil {
		return nil, err
	}
	response := &accommodation_service.AccGetAllResponse{
		Acc: []*accommodation_service.Accommodation{},
	}
	for _, accommodation := range accommodationsLocation {
		current := mapAccommodation(accommodation)
		response.Acc = append(response.Acc, current)
	}
	return response, nil
}

func parseISO8601(dateString string) (time.Time, error) {
	dateTime, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return dateTime, nil
}
func NewUserClient() user_service.UserServiceClient {
	conn, err := grpc.Dial("user-service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return user_service.NewUserServiceClient(conn)
}
