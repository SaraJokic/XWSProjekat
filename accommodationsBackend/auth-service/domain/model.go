package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email"    json:"email"`
	Name     string             `bson:"name"     json:"name"`
	LastName string             `bson:"lastname" json:"lastname"`
	Role     UserType           `bson:"role" json:"role"`
}

type UserType int

const (
	Customer UserType = iota
	Admin
)