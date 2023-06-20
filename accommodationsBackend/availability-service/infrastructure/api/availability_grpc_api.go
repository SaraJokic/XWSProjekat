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
	layout := "2006-01-02 15:04:05 -0700 MST"
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
	availability := mapDomainAvailability(request)

	err := handler.service.Insert(availability)
	if err != nil {
		log.Println("Insert not working")
		return nil, err
	}
	a := availability_service.CreateNewAvailabilityResponse{Poruka: "Successfully created new availability object."}
	return &a, nil
}
func (handler *AvailabilityHandler) UpdateAvailability(ctx context.Context, request *availability_service.UpdateAvailabilityRequest) (*availability_service.UpdateAvailabilityResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	fmt.Println("AVAILABILITY UPDATE: id ", request.Id)
	a := request.Availability
	fmt.Println("AVAILABILITY UPDATE: availability iz requesta ", request.Availability)
	accMapped := mapUpdatevailability(a)
	fmt.Println("AVAILABILITY UPDATE: mapiranooo ", accMapped)
	err = handler.service.Update(objectId, accMapped)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update accommodation")
	}

	return &availability_service.UpdateAvailabilityResponse{Availability: a}, nil
}
func (handler *AvailabilityHandler) UpdateAfterReservation(ctx context.Context, request *availability_service.UpdateAfterReservationRequest) (*availability_service.UpdateAfterReservationResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	availability, _ := handler.service.GetByAccommodationId(objectId)
	updatedSlots := make([]domain.AvailabilitySlot, 0)
	slots := availability.AvailableSlots
	layout := "2006-01-02 15:04:05 -0700 MST"
	startdate, err := time.Parse(layout, request.SelectedStartDate)
	if err != nil {
		log.Println("Failed to parse the string of StartDate: ", err)
		return nil, err
	}
	enddate, err := time.Parse(layout, request.SelectedEndDate)
	if err != nil {
		log.Println("Failed to parse the string of EndDate: ", err)
		return nil, err
	}
	for _, slot := range slots {

		// Check if the slot overlaps with the selected dates
		if slot.StartDate.Before(startdate) && slot.EndDate.After(enddate) {
			// Split the slot into two parts: before the selected start date and after the selected end date
			// Add the first part to the updated slots
			updatedSlots = append(updatedSlots, domain.AvailabilitySlot{SlotId: primitive.NewObjectID(), StartDate: slot.StartDate, EndDate: startdate})

			// Add the second part to the updated slots
			updatedSlots = append(updatedSlots, domain.AvailabilitySlot{SlotId: primitive.NewObjectID(), StartDate: enddate, EndDate: slot.EndDate})
		} else if slot.StartDate.Before(startdate) && slot.EndDate.After(startdate) {
			// Adjust the end date of the slot to be the selected start date
			slot.EndDate = startdate
			updatedSlots = append(updatedSlots, slot)
		} else if slot.StartDate.Before(enddate) && slot.EndDate.After(enddate) {
			// Adjust the start date of the slot to be the selected end date
			slot.StartDate = enddate
			updatedSlots = append(updatedSlots, slot)
		} else if slot.StartDate.Equal(startdate) && slot.EndDate.After(enddate) {

			updatedSlots = append(updatedSlots, domain.AvailabilitySlot{SlotId: primitive.NewObjectID(), StartDate: enddate, EndDate: slot.EndDate})

		} else if slot.StartDate.Before(startdate) && slot.EndDate.Equal(enddate) {

			updatedSlots = append(updatedSlots, domain.AvailabilitySlot{SlotId: primitive.NewObjectID(), StartDate: slot.StartDate, EndDate: startdate})

		} else if slot.StartDate.Equal(startdate) && slot.EndDate.Equal(enddate) {
			continue
		} else {
			// No overlap, keep the slot as it is
			updatedSlots = append(updatedSlots, slot)
		}
	}
	availability.AvailableSlots = updatedSlots
	// Zameniti da su updatedSlots ustv availabilit.AvailabilitySlots
	err = handler.service.Update(availability.Id, availability)
	if err != nil {
		return nil, err
	}
	return &availability_service.UpdateAfterReservationResponse{Message: "Seccessfully changed availability slots"}, nil
}
