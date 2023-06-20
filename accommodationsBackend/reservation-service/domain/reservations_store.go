package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReservationStore interface {
	Get(id primitive.ObjectID) (*Reservation, error)
	GetAll() ([]*Reservation, error)
	Insert(reservation *Reservation) error
	DeleteAll()
	Update(id primitive.ObjectID, reservation *Reservation) error
	GetByAccommodationId(id primitive.ObjectID) ([]*Reservation, error)
	GetByUserId(id primitive.ObjectID) ([]*Reservation, error)
	Delete(id string) error
}
