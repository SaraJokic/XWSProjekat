package persistence

import (
	"accommodationsBackend/common/proto/eventstore"
	"accommodationsBackend/event-store/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "EventStore"
	COLLECTION = "events"
)

type EventsMongoDBStore struct {
	events *mongo.Collection
}

func (e EventsMongoDBStore) CreateEvent(ctx context.Context, event *eventstore.Event) error {
	eventData := bson.M{
		"_id":           event.EventId,
		"eventtype":     event.EventType,
		"aggregateid":   event.AggregateId,
		"aggregatetype": event.AggregateType,
		"eventdata":     event.EventData,
		"stream":        event.Stream,
	}
	_, err := e.events.InsertOne(ctx, eventData)
	if err != nil {
		return fmt.Errorf("error on insert into events: %w", err)
	}
	return nil
}

func (e EventsMongoDBStore) GetEvents(ctx context.Context, filter *eventstore.GetEventsRequest) ([]*eventstore.Event, error) {
	collection := e.events
	query := bson.M{}

	if filter.EventId != "" {
		query["id"] = filter.EventId
	}

	if filter.AggregateId != "" {
		query["aggregateid"] = filter.AggregateId
	}

	cur, err := collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var events []*eventstore.Event
	for cur.Next(ctx) {
		var event eventstore.Event
		err := cur.Decode(&event)
		if err != nil {
			return events, err
		}
		events = append(events, &event)
	}

	if err := cur.Err(); err != nil {
		return events, err
	}

	return events, nil
}

func NewEventsnMongoDBStore(client *mongo.Client) domain.EventStore {
	events := client.Database(DATABASE).Collection(COLLECTION)
	return &EventsMongoDBStore{
		events: events,
	}
}
