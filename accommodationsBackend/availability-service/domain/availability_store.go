package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AvailabilityStore interface {
	Get(id primitive.ObjectID) (*Availability, error)
	GetAll() ([]*Availability, error)
	Insert(availability *Availability) error
	DeleteAll()
	Update(id primitive.ObjectID, availability *Availability) error
	GetByAccommodationId(id primitive.ObjectID) (*Availability, error)
}
