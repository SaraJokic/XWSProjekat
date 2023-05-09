package domain

type Benefits struct {
	Wifi        bool `bson:"wifi"`
	Kitchen     bool `bson:"kitchen"`
	FreeParking bool `bson:"freeparking"`
}
