package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReservationRequest struct {
	Id              primitive.ObjectID `bson:"_id"`
	GuestId         primitive.ObjectID `bson:"guestId"`
	AccommodationId primitive.ObjectID `bson:"accommodationId"`
	StartDate       time.Time          `bson:"startdate"`
	EndDate         time.Time          `bson:"enddate"`
	NmOfGuests      int                `bson:"numofguests"`
}
type Reservation struct {
	Id              primitive.ObjectID `bson:"_id"`
	ApprovedRequest ReservationRequest `bson:"requestId"`
}
