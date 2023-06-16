package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accommodation struct {
	Id               primitive.ObjectID `bson:"_id"`
	HostId           primitive.ObjectID `bson:"hostid"`
	Name             string             `bson:"name"`
	Location         string             `bson:"location"`
	Description      string             `bson:"description"`
	Benefits         Benefits           `bson:"benefits"`
	MinGuests        int                `bson:"minGuests"`
	MaxGuests        int                `bson:"maxGuests"`
	Pictures         []string           `bson:"pictures, omitempty"`
	AutomaticApprove bool               `bson:"automaticapprove, omitempty"`
}
