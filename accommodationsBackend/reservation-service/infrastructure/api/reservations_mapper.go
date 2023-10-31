package api

import (
	"accommodationsBackend/common/proto/reservation_service"
	"accommodationsBackend/reservation-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func mapReservation(r *domain.Reservation) *reservation_service.Reservation {
	reservationMapped := &reservation_service.Reservation{
		Id:              r.Id.Hex(),
		AccommodationId: r.AccommodationId.Hex(),
		GuestId:         r.GuestId.Hex(),
		StartDate:       r.StartDate.String(),
		EndDate:         r.EndDate.String(),
		NumOfGuests:     int64(r.NumOfGuests),
		Status:          reservation_service.Status(r.Status),
		IsCanceled:      r.IsCanceled,
	}
	return reservationMapped
}
func mapDomainReservation(r *reservation_service.CreateReservationRequest) *domain.Reservation {
	accId, err := primitive.ObjectIDFromHex(r.AccommodationId)
	if err != nil {
		return nil
	}
	guestId, err := primitive.ObjectIDFromHex(r.GuestId)
	if err != nil {
		return nil
	}
	layout := "2006-01-02T15:04:05Z"
	//fmt.Println("Start datum price change", request.PriceChange.Startdate)
	startdate, err := time.Parse(layout, r.StartDate)
	if err != nil {
		log.Println("Failed to parse the string of StartDate: ", err)
		return nil
	}
	enddate, err := time.Parse(layout, r.EndDate)
	if err != nil {
		log.Println("Failed to parse the string of EndDate: ", err)
		return nil
	}
	reservationMapped := &domain.Reservation{
		AccommodationId: accId,
		GuestId:         guestId,
		StartDate:       startdate,
		EndDate:         enddate,
		NumOfGuests:     int(r.NumOfGuests),
		Status:          domain.Status(r.Status),
		IsCanceled:      r.Iscanceled,
	}
	return reservationMapped
}
func mapUpdateReservation(r *reservation_service.Reservation) *domain.Reservation {
	accId, err := primitive.ObjectIDFromHex(r.AccommodationId)
	if err != nil {
		return nil
	}
	guestId, err := primitive.ObjectIDFromHex(r.GuestId)
	if err != nil {
		return nil
	}
	rId, err := primitive.ObjectIDFromHex(r.Id)
	if err != nil {
		return nil
	}
	layout := "2006-01-02T15:04:05Z"
	//fmt.Println("Start datum price change", request.PriceChange.Startdate)
	startdate, err := time.Parse(layout, r.StartDate)
	if err != nil {
		log.Println("Failed to parse the string of StartDate: ", err)
		return nil
	}
	enddate, err := time.Parse(layout, r.EndDate)
	if err != nil {
		log.Println("Failed to parse the string of EndDate: ", err)
		return nil
	}
	reservationMapped := &domain.Reservation{
		Id:              rId,
		AccommodationId: accId,
		GuestId:         guestId,
		StartDate:       startdate,
		EndDate:         enddate,
		NumOfGuests:     int(r.NumOfGuests),
		Status:          domain.Status(r.Status),
		IsCanceled:      r.IsCanceled,
	}
	return reservationMapped
}
