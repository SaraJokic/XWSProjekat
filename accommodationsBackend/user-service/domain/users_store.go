package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	Register(user *User) error
	DeleteAll()
	CheckIfEmailAndUsernameExist(email string, username string) (bool, error)
}
