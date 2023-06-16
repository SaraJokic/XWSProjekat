package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Availability struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationId primitive.ObjectID `bson:"accommodationId"`
	AvailableSlots  []AvailabilitySlot `bson:"availability, omitempty"`
	Price           float64            `bson:"price"`
	IsPricePerGuest bool               `bson:"ispriceperperson"`
	ChangePrice     []PriceChange      `bson:"changePrice, omitempty"`
}

type AvailabilitySlot struct {
	SlotId    primitive.ObjectID `bson:"slotId"`
	StartDate time.Time          `bson:"startDate"`
	EndDate   time.Time          `bson:"endDate"`
}

type PriceChange struct {
	StartDate time.Time `bson:"startDate"`
	EndDate   time.Time `bson:"endDate"`
	Change    float64   `bson:"change"`
}
