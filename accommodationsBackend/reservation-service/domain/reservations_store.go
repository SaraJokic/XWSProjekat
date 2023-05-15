package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccommodationStore interface {
	Get(id primitive.ObjectID) (*Reservation, error)
	GetAll() ([]*Reservation, error)
	Insert(reservation *Reservation) error
	DeleteAll()
}
