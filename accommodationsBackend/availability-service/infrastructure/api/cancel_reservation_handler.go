package api

import (
	"accommodationsBackend/availability-service/application"
	"accommodationsBackend/availability-service/domain"
	"accommodationsBackend/common/saga/cancel_reservation"
	saga "accommodationsBackend/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type CancelReservationCommandHandler struct {
	availabilityService *application.AvailabilityService
	replyPublisher      saga.Publisher
	commandSubscriber   saga.Subscriber
}

func NewCancelReservationCommandHandler(availabilityService *application.AvailabilityService, publisher saga.Publisher, subscriber saga.Subscriber) (*CancelReservationCommandHandler, error) {
	o := &CancelReservationCommandHandler{
		availabilityService: availabilityService,
		replyPublisher:      publisher,
		commandSubscriber:   subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}
func (handler *CancelReservationCommandHandler) handle(command *cancel_reservation.CancelReservationCommand) {
	reply := cancel_reservation.CancelReservationReply{Reservation: command.Reservation}
	accid, err := primitive.ObjectIDFromHex(command.Reservation.AccommodationId)
	if err != nil {
		return
	}

	switch command.Type {
	case cancel_reservation.FreeSlotsCommand:
		availability, _ := handler.availabilityService.GetByAccommodationId(accid)

		availability.AvailableSlots = append(availability.AvailableSlots, domain.AvailabilitySlot{SlotId: primitive.NewObjectID(), StartDate: command.Reservation.StartDate, EndDate: command.Reservation.EndDate})
		err = handler.availabilityService.Update(availability.Id, availability)
		if err != nil {
			log.Println("Failed to update Availability Object")
			reply.Type = cancel_reservation.SlotsNotUpdatedReply
		}
		reply.Type = cancel_reservation.SlotsUpdatedReply
	case cancel_reservation.RollbackSlotsCommand:
		//TODO: rollback slotova
		reply.Type = cancel_reservation.SlotsRolledbackReply
	default:
		reply.Type = cancel_reservation.UnknownReply
	}

	if reply.Type != cancel_reservation.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
