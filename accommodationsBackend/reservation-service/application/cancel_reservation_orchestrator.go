package application

import (
	"accommodationsBackend/common/saga/cancel_reservation"
	saga "accommodationsBackend/common/saga/messaging"
	"accommodationsBackend/reservation-service/domain"
)

type CancelReservationOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewCancelReservationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*CancelReservationOrchestrator, error) {
	o := &CancelReservationOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}
func (o *CancelReservationOrchestrator) Start(reservation *domain.Reservation) error {
	event := &cancel_reservation.CancelReservationCommand{
		Type: cancel_reservation.UpdateTimesCancelledCommand,
		Reservation: cancel_reservation.ReservationDetails{
			Id:              reservation.Id.Hex(),
			GuestId:         reservation.GuestId.Hex(),
			AccommodationId: reservation.AccommodationId.Hex(),
			StartDate:       reservation.StartDate,
			EndDate:         reservation.EndDate,
		},
	}
	return o.commandPublisher.Publish(event)
}
func (o *CancelReservationOrchestrator) handle(reply *cancel_reservation.CancelReservationReply) {
	command := cancel_reservation.CancelReservationCommand{
		Reservation: reply.Reservation,
	}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != cancel_reservation.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}
func (o *CancelReservationOrchestrator) nextCommandType(reply cancel_reservation.CancelReservationReplyType) cancel_reservation.CancelReservationCommandType {
	switch reply {
	case cancel_reservation.UserUpdatedReply:
		return cancel_reservation.FreeSlotsCommand
	case cancel_reservation.UserNotUpdatedReply:
		return cancel_reservation.DontCancelReservationCommand
	case cancel_reservation.UserRolledbackReply:
		return cancel_reservation.DontCancelReservationCommand
	case cancel_reservation.SlotsUpdatedReply:
		return cancel_reservation.DeleteReservationCommand
	case cancel_reservation.SlotsNotUpdatedReply:
		return cancel_reservation.RollbackUserUpdateCommand
	case cancel_reservation.ReservationNotCancelledReply:
		return cancel_reservation.RollbackSlotsCommand
	case cancel_reservation.SlotsRolledbackReply:
		return cancel_reservation.RollbackUserUpdateCommand
	default:
		return cancel_reservation.UnknownCommand
	}
}
