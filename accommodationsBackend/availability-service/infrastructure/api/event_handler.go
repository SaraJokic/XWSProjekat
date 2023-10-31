package api

import (
	"accommodationsBackend/availability-service/application"
	"accommodationsBackend/common/proto/reservation_service"
	"accommodationsBackend/common/saga/messaging/nats"
	"encoding/json"
	"fmt"
	nats2 "github.com/nats-io/nats.go"
	"log"
)

type EventHandler struct {
	natsComponent       *nats.NATSComponent
	availabilityService *application.AvailabilityService
	eventCLient         AvailabilityEventClient
}

func NewEventHandler(nc *nats.NATSComponent, availabilityService *application.AvailabilityService) (*EventHandler, error) {
	e := &EventHandler{
		natsComponent:       nc,
		availabilityService: availabilityService,
		eventCLient:         grpcClient{},
	}
	// Creates JetStreamContext for create consumer
	jetStreamContext, err := nc.JetStreamContext()
	if err != nil {
		log.Fatal(err)
	}

	// Create push based consumer as durable
	subscription, err := jetStreamContext.QueueSubscribe("Reservations.Canceled", "availability-service", e.handleMessage,
		nats2.Durable("availability-service"), nats2.ManualAck())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SUBSCRIPTION JE: %v", subscription)
	return e, nil
}
func (h *EventHandler) handleMessage(msg *nats2.Msg) {
	fmt.Println("Procitao poruku !!!!!!!")
	msg.Ack()
	var reservation reservation_service.Reservation
	err := json.Unmarshal(msg.Data, &reservation)
	if err != nil {
		log.Print(err)
		return
	}
	availability, _ := h.availabilityService.MakeSlotAvailable(reservation.AccommodationId, reservation.StartDate, reservation.EndDate)
	//h.availabilityService.MakeSlotAvailable(reservation.AccommodationId, reservation.StartDate, reservation.EndDate)
	h.eventCLient.newAvailableSlot(availability)

	log.Printf("Message subscribed on subject:%s, from:%s, data:%s",
		"Accommodation.Created", "availability-service", msg.Data)
}
