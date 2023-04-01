package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId   string             `bson:"userid,omitempty" json:"userid"`
	FlightId string             `bson:"flightid,omitempty" json:"flightid"`
	Quantity int                `bson:"quantity,omitempty" json:"quantity"`
	Expired  bool               `bson:"expired,omitempty" json:"expired"`
}
type Tickets []*Ticket

func (t *Tickets) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func (t *Ticket) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)
}

func (t *Ticket) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(t)
}
