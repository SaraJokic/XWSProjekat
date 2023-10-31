package api

import (
	"accommodationsBackend/common/proto/eventstore"
	"accommodationsBackend/reservation-service/domain"
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

type ReservationEventClient interface {
	cancelReservation(reservation domain.Reservation) error
	createReservation(reservation domain.Reservation) error
}
type grpcClient struct {
}

func (gc grpcClient) cancelReservation(reservation domain.Reservation) error {
	fmt.Println("Usao u cancle reservation event client za upisivanje eventa")
	conn, err := grpc.Dial("event-store:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()
	client := eventstore.NewEventStoreClient(conn)
	eventid, _ := uuid.NewUUID()
	reservationJSON, _ := json.Marshal(reservation)
	event := &eventstore.Event{
		EventId:       eventid.String(),
		EventType:     "Reservations.Canceled",
		EventTime:     time.Now().Format("2006-01-02 15:04:05"),
		AggregateId:   reservation.Id.Hex(),
		AggregateType: "Reservation",
		EventData:     string(reservationJSON),
		Stream:        "Reservations",
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

func (gc grpcClient) createReservation(reservation domain.Reservation) error {
	conn, err := grpc.Dial("event-store:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()
	client := eventstore.NewEventStoreClient(conn)
	eventid, _ := uuid.NewUUID()
	reservationJSON, _ := json.Marshal(reservation)
	event := &eventstore.Event{
		EventId:       eventid.String(),
		EventType:     "Reservations.Created",
		AggregateId:   reservation.Id.Hex(),
		AggregateType: "Reservation",
		EventData:     string(reservationJSON),
		Stream:        "Reservations",
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
