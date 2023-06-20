package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Status int

const (
	Pending  Status = iota // 0
	Approved               // 1
	Denied                 // 2
)

type Reservation struct {
	Id              primitive.ObjectID `bson:"_id"`
	GuestId         primitive.ObjectID `bson:"guestId"`
	AccommodationId primitive.ObjectID `bson:"accommodationId"`
	StartDate       time.Time          `bson:"startdate"`
	EndDate         time.Time          `bson:"enddate"`
	NumOfGuests     int                `bson:"numofguests"`
	Status          Status             `bson:"status"`
}
