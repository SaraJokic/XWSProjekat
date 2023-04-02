package model

import (
	"encoding/json"
	"io"
	"time"
)

type FlightSearchDTO struct {
	FromPlace        string    `bson:"fromplace" json:"fromplace" gorm:"not null"`
	ToPlace          string    `bson:"toplace" json:"toplace" gorm:"not null"`
	StartTime        time.Time `bson:"starttime" json:"starttime" gorm:"not null"`
	RemainingTickets int64     `bson:"remainingtickets" json:"remainingtickets" gorm:"not null"`
}

func (f *FlightSearchDTO) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(f)
}

func (f *FlightSearchDTO) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(f)
}
