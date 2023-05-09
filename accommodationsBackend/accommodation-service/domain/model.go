package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accommodation struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Location  string             `bson:"location"`
	Benefits  Benefits           `bson:"benefits"`
	MinGuests int                `bson:"minGuests"`
	MaxGuests int                `bson:"maxGuests"`
}
