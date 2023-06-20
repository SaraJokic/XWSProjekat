package api

import (
	"accommodationsBackend/common/saga/cancel_reservation"
	saga "accommodationsBackend/common/saga/messaging"
	"accommodationsBackend/reservation-service/application"
)

type CancelReservationCommandHandler struct {
	reservationService *application.ReservationService
	replyPublisher     saga.Publisher
	commandSubscriber  saga.Subscriber
}

func NewCancelReservationCommandHandler(reservationService *application.ReservationService, publisher saga.Publisher, subscriber saga.Subscriber) (*CancelReservationCommandHandler, error) {
	o := &CancelReservationCommandHandler{
		reservationService: reservationService,
		replyPublisher:     publisher,
		commandSubscriber:  subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}
func (handler *CancelReservationCommandHandler) handle(command *cancel_reservation.CancelReservationCommand) {
	reply := cancel_reservation.CancelReservationReply{Reservation: command.Reservation}

	switch command.Type {
	case cancel_reservation.DeleteReservationCommand:
		err := handler.reservationService.Cancel(command.Reservation.Id)
		if err != nil {
			reply.Type = cancel_reservation.ReservationNotCancelledReply
			return
		}
		reply.Type = cancel_reservation.CancelledReservationReply
	case cancel_reservation.DontCancelReservationCommand:
		reply.Type = cancel_reservation.ReservationNotCancelledReply
	default:
		reply.Type = cancel_reservation.UnknownReply
	}

	if reply.Type != cancel_reservation.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
