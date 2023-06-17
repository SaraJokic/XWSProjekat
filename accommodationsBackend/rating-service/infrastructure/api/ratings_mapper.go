package api

import (
	"accommodationsBackend/common/proto/rating_service"
	"accommodationsBackend/rating-service/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func mapAccommodationRating(acc *domain.RateAccommodation) *rating_service.RateAccommodation {
	mapped := &rating_service.RateAccommodation{
		Id:              acc.Id.Hex(),
		GuestId:         acc.GuestId.Hex(),
		DateRating:      acc.Date.String(),
		Rating:          float32(acc.Rating),
		AccommodationId: acc.AccommodationId.Hex(),
	}
	return mapped
}
func mapHostRating(acc *domain.RateHost) *rating_service.RateHost {
	mapped := &rating_service.RateHost{
		Id:         acc.Id.Hex(),
		GuestId:    acc.GuestId.Hex(),
		DateRating: acc.Date.String(),
		Rating:     float32(acc.Rating),
		HostId:     acc.HostId.Hex(),
	}
	return mapped
}
func mapNewHostRating(rating *rating_service.CreateNewHostRatingRequest) *domain.RateHost {
	hostId, err := primitive.ObjectIDFromHex(rating.HostId)
	guestId, err := primitive.ObjectIDFromHex(rating.GuestId)
	layout := "2006-01-02T15:04:05Z"
	newdate, err := time.Parse(layout, rating.DateRating)
	if err != nil {
		log.Println("Failed to parse the string of EndDate: ", err)
		return nil
	}
	fmt.Println("-------------------RATING MAPPING Ovo je guest id koji se cuva: ", guestId)
	fmt.Println("-------------------RATING MAPPING Ovo je host id koji se cuva: ", hostId)
	mapped := &domain.RateHost{
		GuestId: guestId,
		HostId:  hostId,
		Rating:  float64(rating.Rating),
		Date:    newdate,
	}
	return mapped
}
func mapNewAccommodationRating(rating *rating_service.CreateNewAccommodationRatingRequest) *domain.RateAccommodation {
	hostId, err := primitive.ObjectIDFromHex(rating.AccommodationId)
	guestId, err := primitive.ObjectIDFromHex(rating.GuestId)
	layout := "2006-01-02T15:04:05Z"
	newdate, err := time.Parse(layout, rating.DateRating)
	if err != nil {
		log.Println("Failed to parse the string of EndDate: ", err)
		return nil
	}
	mapped := &domain.RateAccommodation{
		GuestId:         guestId,
		AccommodationId: hostId,
		Rating:          float64(rating.Rating),
		Date:            newdate,
	}
	return mapped
}
