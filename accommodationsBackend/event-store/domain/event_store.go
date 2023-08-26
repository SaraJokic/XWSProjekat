package domain

import (
	"accommodationsBackend/common/proto/eventstore"
	"context"
)

type EventStore interface {
	CreateEvent(ctx context.Context, event *eventstore.Event) error
	GetEvents(ctx context.Context, filter *eventstore.GetEventsRequest) ([]*eventstore.Event, error)
}
