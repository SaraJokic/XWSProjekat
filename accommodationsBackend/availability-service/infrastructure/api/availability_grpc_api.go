package api

import (
	"accommodationsBackend/availability-service/application"
	"accommodationsBackend/availability-service/domain"
	availability_service "accommodationsBackend/common/proto/availability-service"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type AvailabilityHandler struct {
	availability_service.UnimplementedAvailabilityServiceServer
	service *application.AvailabilityService
}

func NewAvailabilityHandler(service *application.AvailabilityService) *AvailabilityHandler {
	return &AvailabilityHandler{
		service: service,
	}
}

func (handler *AvailabilityHandler) Get(ctx context.Context, request *availability_service.GetAvailableRequest) (*availability_service.GetAvailableResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	availability, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	aMapped := mapAvailability(availability)
	response := &availability_service.GetAvailableResponse{
		Availability: aMapped,
	}
	return response, nil
}

func (handler *AvailabilityHandler) GetAll(ctx context.Context, request *availability_service.GetAllAvailableRequest) (*availability_service.GetAllAvailableResponse, error) {
	availabilities, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &availability_service.GetAllAvailableResponse{
		Availability: []*availability_service.Availability{},
	}
	for _, a := range availabilities {
		current := mapAvailability(a)
		response.Availability = append(response.Availability, current)
	}
	return response, nil
}
func (handler *AvailabilityHandler) CreatePriceChange(ctx context.Context, request *availability_service.CreatePriceChangeRequest) (*availability_service.CreatePriceChangeResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.AvailabilityID)
	if err != nil {
		return nil, err
	}
	availability, err := handler.service.Get(id)
	if err != nil {
		log.Println("Failed to get availability object from the database:", err)
		return nil, status.Errorf(codes.Internal, "Failed to get availability object from the database")
	}
	layout := "2006-01-02T15:04:05Z"
	//fmt.Println("Start datum price change", request.PriceChange.Startdate)
	startdate, err := time.Parse(layout, request.PriceChange.Startdate)
	if err != nil {
		log.Println("Failed to parse the string of StartDate: ", err)
		return nil, err
	}
	enddate, err := time.Parse(layout, request.PriceChange.Enddate)
	if err != nil {
		log.Println("Failed to parse the string of EndDate: ", err)
		return nil, err
	}
	availability.ChangePrice = append(availability.ChangePrice, domain.PriceChange{StartDate: startdate, EndDate: enddate, Change: float64(request.PriceChange.Change)})
	//fmt.Println(availability.ChangePrice)
	err = handler.service.Update(id, availability)
	if err != nil {
		log.Println("Failed to update Availability Object")
		return nil, err
	}
	response := &availability_service.CreatePriceChangeResponse{
		Poruka: "PriceChange added successfully",
	}
	return response, nil
}

func (handler *AvailabilityHandler) AddAvailableSlot(ctx context.Context, request *availability_service.AddAvailableSlotRequest) (*availability_service.AddAvailableSlotResponse, error) {
	fmt.Println("Ovo je ceo availability request:", request)
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	availability, err := handler.service.Get(id)
	if err != nil {
		log.Println("Failed to get availability object from the database:", err)
		return nil, status.Errorf(codes.Internal, "Failed to get availability object from the database")
	}
	layout := "2006-01-02T15:04:05Z"
	//fmt.Println("Start datum price change", request.PriceChange.Startdate)
	startdate, err := time.Parse(layout, request.AvailableSlot.StartDate)
	if err != nil {
		log.Println("Failed to parse the string of StartDate: ", err)
		return nil, err
	}
	enddate, err := time.Parse(layout, request.AvailableSlot.EndDate)
	if err != nil {
		log.Println("Failed to parse the string of EndDate: ", err)
		return nil, err
	}
	availability.AvailableSlots = append(availability.AvailableSlots, domain.AvailabilitySlot{SlotId: primitive.NewObjectID(), StartDate: startdate, EndDate: enddate})
	//fmt.Println(availability.ChangePrice)
	err = handler.service.Update(id, availability)
	if err != nil {
		log.Println("Failed to update Availability Object")
		return nil, err
	}
	response := &availability_service.AddAvailableSlotResponse{Poruka: "New available slot added successfully."}

	return response, nil
}

func (handler *AvailabilityHandler) GetByAccommodationId(ctx context.Context, request *availability_service.GetByAccIdRequest) (*availability_service.GetByAccIdResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	availability, err := handler.service.GetByAccommodationId(objectId)
	if err != nil {
		return nil, err
	}
	aMapped := mapAvailability(availability)
	response := &availability_service.GetByAccIdResponse{
		Availability: aMapped,
	}
	return response, nil
}
func (handler *AvailabilityHandler) CreateNewAvailability(ctx context.Context, request *availability_service.CreateNewAvailabilityRequest) (*availability_service.CreateNewAvailabilityResponse, error) {
	fmt.Println("Ovo je ceo availability request:", request)
	availability := mapDomainAvailability(request)

	err := handler.service.Insert(availability)
	if err != nil {
		log.Println("Insert not working")
		return nil, err
	}
	a := availability_service.CreateNewAvailabilityResponse{Poruka: "Successfully created new availability object."}
	return &a, nil
}
