package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username       string             `bson:"username" json:"username"`
	Password       string             `bson:"password" json:"password"`
	Email          string             `bson:"email"    json:"email"`
	Name           string             `bson:"name"     json:"name"`
	LastName       string             `bson:"lastname" json:"lastname"`
	City           string             `bson:"city"     json:"city"`
	Country        string             `bson:"country"  json:"country"`
	Role           UserType           `bson:"role" json:"role"`
	TimesCancelled int                `bson:"timescancelled" json:"timescancelled"`
	ProminentHost  bool               `bson:"prominentHost"   json:"prominentHost"`
}

type UserType int32

const (
	Customer UserType = iota
	Admin
)
