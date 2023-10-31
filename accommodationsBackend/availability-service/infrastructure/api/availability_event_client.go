package api

import (
	"accommodationsBackend/availability-service/domain"
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
	"time"
)

type AvailabilityEventClient interface {
	newAvailableSlot(acc domain.Availability) error
}
type grpcClient struct {
}

func (gc grpcClient) newAvailableSlot(availability domain.Availability) error {
	conn, err := grpc.Dial("event-store:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()
	client := eventstore.NewEventStoreClient(conn)
	eventid, _ := uuid.NewUUID()
	availabilityJSON, _ := json.Marshal(availability)
	event := &eventstore.Event{
		EventId:       eventid.String(),
		EventType:     "Availability.SlotAdded",
		EventTime:     time.Now().Format("2006-01-02 15:04:05"),
		AggregateId:   availability.Id.Hex(),
		AggregateType: "Availability",
		EventData:     string(availabilityJSON),
		Stream:        "Availabilities",
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
