package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RateHost struct {
	Id      primitive.ObjectID `bson:"_id"`
	GuestId primitive.ObjectID `bson:"guestId"`
	Date    time.Time          `bson:"dateRating"`
	Rating  float64            `bson:"rating"`
	HostId  primitive.ObjectID `bson:"hostId"`
}
type RateAccommodation struct {
	Id              primitive.ObjectID `bson:"_id"`
	GuestId         primitive.ObjectID `bson:"guestId"`
	Date            time.Time          `bson:"dateRating"`
	Rating          float64            `bson:"rating"`
	AccommodationId primitive.ObjectID `bson:"accommodationId"`
}
