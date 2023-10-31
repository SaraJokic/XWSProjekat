package api

import (
	availability_service "accommodationsBackend/common/proto/availability-service"
	"accommodationsBackend/common/proto/reservation_service"
	"accommodationsBackend/reservation-service/application"
	"accommodationsBackend/reservation-service/domain"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type ReservationHandler struct {
	reservation_service.UnimplementedReservationServiceServer
	service *application.ReservationService
	rpc     ReservationEventClient
}

func NewReservationHandler(service *application.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		service: service,
		rpc:     grpcClient{},
	}
}
func (handler *ReservationHandler) Get(ctx context.Context, request *reservation_service.GetReservationRequest) (*reservation_service.GetReservationResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	reservation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	rMapped := mapReservation(reservation)
	response := &reservation_service.GetReservationResponse{
		Reservation: rMapped,
	}
	return response, nil
}

func (handler *ReservationHandler) GetAll(ctx context.Context, request *reservation_service.GetAllReservationsRequest) (*reservation_service.GetAllReservationsResponse, error) {
	reservations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &reservation_service.GetAllReservationsResponse{
		Reservations: []*reservation_service.Reservation{},
	}
	for _, r := range reservations {
		current := mapReservation(r)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) CreateNewReservation(ctx context.Context, request *reservation_service.CreateReservationRequest) (*reservation_service.CreateReservationResponse, error) {
	reservation := mapDomainReservation(request)
	if handler.isReservationAllowed(reservation) {
		err := handler.service.Insert(reservation)
		if err != nil {
			log.Println("Insert not working")
			return nil, err
		}
		r := reservation_service.CreateReservationResponse{Id: reservation.Id.Hex()}
		return &r, nil
	}
	return nil, errors.New("You cannot make a reservation in that period. It's not available.")

}
func (handler *ReservationHandler) UpdateReservation(ctx context.Context, request *reservation_service.UpdateReservationRequest) (*reservation_service.UpdateReservationResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	r := request.Reservation
	reservationMapped := mapUpdateReservation(r)
	err = handler.service.Update(objectId, reservationMapped) //menja se polje isCancelled na true
	err = handler.rpc.cancelReservation(*reservationMapped)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update reservation")
	}

	return &reservation_service.UpdateReservationResponse{Reservation: r}, nil
}

func (handler *ReservationHandler) GetByAccommodationId(ctx context.Context, request *reservation_service.GetReservationByAccIdRequest) (*reservation_service.GetAllReservationsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	reservations, err := handler.service.GetByAccommodationId(objectId)
	if err != nil {
		return nil, err
	}
	response := &reservation_service.GetAllReservationsResponse{
		Reservations: []*reservation_service.Reservation{},
	}
	for _, r := range reservations {
		current := mapReservation(r)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) GetReservationByUserId(ctx context.Context, request *reservation_service.GetReservationByUserIdRequest) (*reservation_service.GetAllReservationsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	reservations, err := handler.service.GetByUserId(objectId)
	if err != nil {
		return nil, err
	}
	response := &reservation_service.GetAllReservationsResponse{
		Reservations: []*reservation_service.Reservation{},
	}
	for _, r := range reservations {
		current := mapReservation(r)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) ChangeStatusReservation(ctx context.Context, request *reservation_service.ChangeStatusReservationRequest) (*reservation_service.ChangeStatusReservationResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	reservation, err := handler.service.Get(id)
	if err != nil {
		log.Println("Failed to get reservation object from the database:", err)
		return nil, status.Errorf(codes.Internal, "Failed to get availability object from the database")
	}
	if request.Status == 1 {
		client := NewAvailabilityClient()
		client.UpdateAfterReservation(context.Background(), &availability_service.UpdateAfterReservationRequest{Id: reservation.AccommodationId.Hex(), SelectedStartDate: reservation.StartDate.String(), SelectedEndDate: reservation.EndDate.String()})
		handler.DenyOtherRequests(reservation)
	}
	reservation.Status = domain.Status(request.Status)
	fmt.Println("Status promenjen na: ", reservation.Status)
	err = handler.service.Update(id, reservation)
	if err != nil {
		log.Println("Failed to update Reservation Object")
		return nil, err
	}
	return &reservation_service.ChangeStatusReservationResponse{Message: "Successfully updated status."}, nil
}
func (handler *ReservationHandler) DeleteReservation(ctx context.Context, request *reservation_service.DeleteReservationRequest) (*reservation_service.DeleteReservationResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	err = handler.service.Cancel(id)
	if err != nil {
		return &reservation_service.DeleteReservationResponse{Message: "Reservation cancel failed"}, nil
	}
	reservation, _ := handler.service.Get(objectId)
	handler.rpc.cancelReservation(*reservation)
	return &reservation_service.DeleteReservationResponse{Message: "Reservation canceled"}, nil
}
func NewAvailabilityClient() availability_service.AvailabilityServiceClient {
	conn, err := grpc.Dial("availability-service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Availability service: %v", err)
	}
	return availability_service.NewAvailabilityServiceClient(conn)
}
func (handler *ReservationHandler) DenyOtherRequests(r *domain.Reservation) {
	reservations, _ := handler.service.GetByAccommodationId(r.AccommodationId)
	for _, reservation := range reservations {
		if handler.checkDateOverlap(reservation, r) {
			handler.ChangeStatusReservation(context.Background(), &reservation_service.ChangeStatusReservationRequest{Id: reservation.Id.Hex(), Status: 2})
		}
	}
}
func (handler *ReservationHandler) checkDateOverlap(r *domain.Reservation, approvedReservation *domain.Reservation) bool {
	if approvedReservation.StartDate.Before(r.EndDate) && approvedReservation.EndDate.After(r.StartDate) {
		return true // preklapaju se
	}
	return false //ne preklapaju se
}
func (handler *ReservationHandler) isReservationAllowed(r *domain.Reservation) bool {
	client := NewAvailabilityClient()
	availability, _ := client.GetByAccommodationId(context.Background(), &availability_service.GetByAccIdRequest{Id: r.AccommodationId.Hex()})
	layout := "2006-01-02 15:04:05 -0700 MST"
	for _, slot := range availability.Availability.AvailableSlots {
		startdate, err := time.Parse(layout, slot.StartDate)
		if err != nil {
			log.Println("Failed to parse the string of StartDate: ", err)
			return false
		}
		enddate, err := time.Parse(layout, slot.EndDate)
		if err != nil {
			log.Println("Failed to parse the string of EndDate: ", err)
			return false
		}
		startdate = time.Date(startdate.Year(), startdate.Month(), startdate.Day(), 0, 0, 0, 0, startdate.Location())
		enddate = time.Date(enddate.Year(), enddate.Month(), enddate.Day(), 0, 0, 0, 0, enddate.Location())
		if (r.StartDate.After(startdate) || r.StartDate.Equal(startdate)) && (r.EndDate.Before(enddate) || r.EndDate.Equal(enddate)) {
			return true
		}
	}
	return false
}
func convertDateTime(dateTimeStr string) (string, error) {
	layoutIn := "2006-01-02 15:04:05 -0700 MST"  // Input layout
	layoutOut := "2006-01-02 00:00:00 -0700 MST" // Output layout

	// Parse input date and time
	t, err := time.Parse(layoutIn, dateTimeStr)
	if err != nil {
		return "", err
	}

	// Format output date and time
	formattedDateTime := t.Format(layoutOut)

	return formattedDateTime, nil
}
func (handler *ReservationHandler) GetReservationByHostId(ctx context.Context, request *reservation_service.GetReservationByUserIdRequest) (*reservation_service.GetAllReservationsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	reservations, err := handler.service.GetByHostId(objectId)
	if err != nil {
		return nil, err
	}
	response := &reservation_service.GetAllReservationsResponse{
		Reservations: []*reservation_service.Reservation{},
	}
	for _, r := range reservations {
		current := mapReservation(r)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
