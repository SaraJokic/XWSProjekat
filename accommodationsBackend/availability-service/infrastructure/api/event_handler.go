package api

import (
	"accommodationsBackend/common/saga/messaging/nats"
	"fmt"
	nats2 "github.com/nats-io/nats.go"
	"log"
)

type EventHandler struct {
	natsComponent *nats.NATSComponent
}

func NewEventHandler(nc *nats.NATSComponent) (*EventHandler, error) {
	e := &EventHandler{
		natsComponent: nc,
	}
	// Creates JetStreamContext for create consumer
	jetStreamContext, err := nc.JetStreamContext()
	if err != nil {
		log.Fatal(err)
	}

	// Create push based consumer as durable
	jetStreamContext.QueueSubscribe("Accommodations.Created", "availability-service", e.handleMessage,
		nats2.Durable("availability-service"), nats2.ManualAck())
	return e, nil
}
func (h *EventHandler) handleMessage(msg *nats2.Msg) {
	msg.Ack()

	log.Printf("Message subscribed on subject:%s, from:%s, data:%s",
		"Accommodation.Created", "availability-service", msg.Data)

	fmt.Printf("Message subscribed on subject:%s, from:%s, data:%v\n",
		"Accommodation.Created", "availability-service", msg.Data)
}
