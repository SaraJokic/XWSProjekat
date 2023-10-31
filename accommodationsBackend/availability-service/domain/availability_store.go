package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AvailabilityStore interface {
	Get(id primitive.ObjectID) (*Availability, error)
	GetAll() ([]*Availability, error)
	Insert(availability *Availability) error
	DeleteAll()
	Update(id primitive.ObjectID, availability *Availability) error
	GetByAccommodationId(id primitive.ObjectID) (*Availability, error)
	FindAvailabilitySlotsByDateRange(startDate time.Time, endDate time.Time) ([]AvailabilitySlot, error)
	MakeSlotAvailable(id primitive.ObjectID, startDate time.Time, endDate time.Time) (*Availability, error)
}
