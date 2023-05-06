package domain

import "github.com/google/uuid"

type Benefits struct {
	ID            uuid.UUID `bson:"id"`
	Wifi          bool      `bson:"wifi"`
	Kitchen       bool      `bson:"kitchen"`
	FreeParking   bool      `bson:"freeParking"`
	Accommodation Accommodation
}
