package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FromPlace   string             `bson:"fromplace" json:"fromplace"`
	ToPlace     string             `bson:"toplace" json:"toplace"`
	StartTime   time.Time          `bson:"starttime,omitempty" json:"starttime"`
	EndTime     time.Time          `bson:"endtime,omitempty" json:"endtime"`
	TicketPrice float64            `bson:"ticketprice,omitempty" json:"ticketprice"`
	NumOfSeats  int                `bson:"numofseats,omitempty" json:"numofseats"`
	TotalSum    float64            `bson:"totalsum,omitempty" json:"totalsum"`
}

type Flights []*Flight

func (f *Flights) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(f)
}

func (f *Flights) FromJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(f)
}

func (p *Flight) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Flight) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
