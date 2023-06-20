package api

import (
	"accommodationsBackend/common/saga/cancel_reservation"
	saga "accommodationsBackend/common/saga/messaging"
	"accommodationsBackend/user-service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CancelReservationCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCancelReservationCommandHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*CancelReservationCommandHandler, error) {
	o := &CancelReservationCommandHandler{
		userService:       userService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}
func (handler *CancelReservationCommandHandler) handle(command *cancel_reservation.CancelReservationCommand) {
	reply := cancel_reservation.CancelReservationReply{Reservation: command.Reservation}
	userId := command.Reservation.GuestId
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return
	}
	u, err := handler.userService.Get(id)
	u.TimesCancelled = u.TimesCancelled + 1

	switch command.Type {
	case cancel_reservation.UpdateTimesCancelledCommand:
		err := handler.userService.UpdateUser(userId, u)
		if err != nil {
			return
		}
		reply.Type = cancel_reservation.UserUpdatedReply
	case cancel_reservation.RollbackUserUpdateCommand:
		u.TimesCancelled = u.TimesCancelled - 1
		err := handler.userService.UpdateUser(userId, u)
		if err != nil {
			return
		}
		reply.Type = cancel_reservation.UserRolledbackReply
	default:
		reply.Type = cancel_reservation.UnknownReply
	}

	if reply.Type != cancel_reservation.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
