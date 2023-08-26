package api

import (
	"accommodationsBackend/common/proto/eventstore"
	"accommodationsBackend/common/saga/messaging/nats"
	"accommodationsBackend/event-store/domain"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// publishEvent publishes an event via NATS JetStream server
func publishEvent(component *nats.NATSComponent, event *eventstore.Event) {
	// Creates JetStreamContext to publish messages into JetStream Stream
	jetStreamContext, _ := component.JetStreamContext()
	subject := event.EventType
	eventMsg := []byte(event.EventData)
	// Publish message on subject (channel)
	jetStreamContext.Publish(subject, eventMsg)
	log.Println("Published message on subject: " + subject)
}

type EventStoreHandler struct {
	eventstore.UnimplementedEventStoreServer
	store domain.EventStore
	nats  *nats.NATSComponent
}

func NewEventStoreHandler(store domain.EventStore, nats *nats.NATSComponent) *EventStoreHandler {
	return &EventStoreHandler{
		store: store,
		nats:  nats,
	}
}

// CreateEvent creates a new event into the event store
func (s *EventStoreHandler) CreateEvent(ctx context.Context, eventRequest *eventstore.CreateEventRequest) (*eventstore.CreateEventResponse, error) {
	err := s.store.CreateEvent(ctx, eventRequest.Event)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	log.Println("Event is created")
	go publishEvent(s.nats, eventRequest.Event)
	return &eventstore.CreateEventResponse{IsSuccess: true, Error: ""}, nil
}

// GetEvents gets all events for the given aggregate and event
func (s *EventStoreHandler) GetEvents(ctx context.Context, filter *eventstore.GetEventsRequest) (*eventstore.GetEventsResponse, error) {
	events, err := s.store.GetEvents(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &eventstore.GetEventsResponse{Events: events}, nil
}

// GetEventsStream get stream of events for the given event
func (s *EventStoreHandler) GetEventsStream(*eventstore.GetEventsRequest, eventstore.EventStore_GetEventsStreamServer) error {
	return nil
}
