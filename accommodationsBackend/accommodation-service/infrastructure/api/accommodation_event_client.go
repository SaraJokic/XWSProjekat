package api

import (
	"accommodationsBackend/accommodations-service/domain"
	"accommodationsBackend/common/proto/eventstore"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
)

type AccommodationEventClient interface {
	createAccommodation(acc domain.Accommodation) error
}
type grpcClient struct {
}

func (gc grpcClient) createAccommodation(acc domain.Accommodation) error {
	conn, err := grpc.Dial("event-store:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()
	client := eventstore.NewEventStoreClient(conn)
	eventid, _ := uuid.NewUUID()
	accommodationJSON, _ := json.Marshal(acc)
	event := &eventstore.Event{
		EventId:       eventid.String(),
		EventType:     "Accommodations.Created",
		AggregateId:   acc.Id.Hex(),
		AggregateType: "Accommodation",
		EventData:     string(accommodationJSON),
		Stream:        "Accommodations",
	}
	createEventRequest := &eventstore.CreateEventRequest{Event: event}
	response, err := client.CreateEvent(context.Background(), createEventRequest)

	if err != nil {
		if st, ok := status.FromError(err); ok {
			fmt.Printf("gRPC status code: %s, message: %s\n", st.Code().String(), st.Message())
			return fmt.Errorf("error from RPC server with: status code:%s message:%s", st.Code().String(), st.Message())
		}
		return fmt.Errorf("error from RPC server: %w", err)
	}
	if response.IsSuccess {
		return nil
	}
	return errors.New("error from RPC server")
}
