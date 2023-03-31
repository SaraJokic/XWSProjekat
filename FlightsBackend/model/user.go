package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email"    json:"email"`
	Name     string             `bson:"name"     json:"name"`
	LastName string             `bson:"lastname" json:"lastname"`
	Role     UserType           `bson:"role" json:"role"`
}

type Users []*User

func (f *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(f)
}

func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
