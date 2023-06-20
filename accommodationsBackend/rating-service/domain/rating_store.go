package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RatingsStore interface {
	GetRateHost(id primitive.ObjectID) (*RateHost, error)
	GetRateAccommodation(id primitive.ObjectID) (*RateAccommodation, error)
	GetAccommodationsRatingsByGuestId(id primitive.ObjectID) ([]*RateAccommodation, error)
	GetHostRatingsByGuestId(id primitive.ObjectID) ([]*RateHost, error)
	GetAccommodationsRatingsByAccommodationId(id primitive.ObjectID) ([]*RateAccommodation, error)
	GetHostRatingsByHostId(id primitive.ObjectID) ([]*RateHost, error)

	InsertAccommodationRating(ratingAccommodation *RateAccommodation) error
	InsertHostRating(ratingHost *RateHost) error

	UpdateAccommodationRating(id string, rating *RateAccommodation) error
	UpdateHostRating(id string, rating *RateHost) error
	DeleteAll()
	DeleteRateHost(id primitive.ObjectID) error
	DeleteRateAccommodation(id primitive.ObjectID) error
}
