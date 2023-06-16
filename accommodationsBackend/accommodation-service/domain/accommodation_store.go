package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AccommodationStore interface {
	Get(id primitive.ObjectID) (*Accommodation, error)
	GetAll() ([]*Accommodation, error)
	Insert(accommodation *Accommodation) error
	DeleteAll()
	UpdateAccommodation(id string, user *Accommodation) error
	GetAccommodationByUserId(id primitive.ObjectID) ([]*Accommodation, error)
	GetAccommodationByLocation(location string) ([]*Accommodation, error)
	GetAccommodationByNumberOfGuests(guest int) ([]*Accommodation, error)
	GetAccommodationByDateRange(startDate, endDate time.Time) ([]*Accommodation, error)
}
