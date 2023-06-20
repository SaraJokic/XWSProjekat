package cancel_reservation

import "time"

type ReservationDetails struct {
	Id              string
	GuestId         string
	AccommodationId string
	StartDate       time.Time
	EndDate         time.Time
}

type CancelReservationCommandType int8

const (
	UpdateTimesCancelledCommand CancelReservationCommandType = iota
	RollbackUserUpdateCommand
	FreeSlotsCommand
	RollbackSlotsCommand
	DeleteReservationCommand
	DontCancelReservationCommand
	UnknownCommand
)

type CancelReservationCommand struct {
	Reservation ReservationDetails
	Type        CancelReservationCommandType
}

type CancelReservationReplyType int8

const (
	UserUpdatedReply CancelReservationReplyType = iota
	UserNotUpdatedReply
	UserRolledbackReply
	SlotsUpdatedReply
	SlotsRolledbackReply
	SlotsNotUpdatedReply
	CancelledReservationReply
	ReservationNotCancelledReply
	UnknownReply
)

type CancelReservationReply struct {
	Reservation ReservationDetails
	Type        CancelReservationReplyType
}
