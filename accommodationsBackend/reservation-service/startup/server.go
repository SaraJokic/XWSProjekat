package startup

import (
	"accommodationsBackend/common/proto/reservation_service"
	saga "accommodationsBackend/common/saga/messaging"
	"accommodationsBackend/common/saga/messaging/nats"
	"accommodationsBackend/reservation-service/application"
	"accommodationsBackend/reservation-service/domain"
	"accommodationsBackend/reservation-service/infrastructure/api"
	"accommodationsBackend/reservation-service/infrastructure/persistence"
	"accommodationsBackend/reservation-service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "reservation_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	reservationsStore := server.initReservationStore(mongoClient)

	commandPublisher := server.initPublisher(server.config.CancelReservationCommandSubject)
	replySubscriber := server.initSubscriber(server.config.CancelReservationReplySubject, QueueGroup)
	cancelReservationOrchestrator := server.initCancelReservationOrchestrator(commandPublisher, replySubscriber)

	reservationsService := server.initReservationService(reservationsStore, cancelReservationOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.CancelReservationCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CancelReservationReplySubject)
	server.initCancelReservationHandler(reservationsService, replyPublisher, commandSubscriber)

	reservationsHandler := server.initReservationHandler(reservationsService)

	server.startGrpcServer(reservationsHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ReservationDBHost, server.config.ReservationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initReservationStore(client *mongo.Client) domain.ReservationStore {
	store := persistence.NewReservationMongoDBStore(client)
	store.DeleteAll()
	/*for _, a := range reservations {
		err := store.Insert(a)
		if err != nil {
			log.Fatal(err)
		}
	}*/
	return store
}

func (server *Server) initReservationService(store domain.ReservationStore, orchestrator *application.CancelReservationOrchestrator) *application.ReservationService {
	return application.NewreservationService(store, orchestrator)
}

func (server *Server) initReservationHandler(service *application.ReservationService) *api.ReservationHandler {
	return api.NewReservationHandler(service)
}
func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}
func (server *Server) initCancelReservationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.CancelReservationOrchestrator {
	orchestrator, err := application.NewCancelReservationOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}
func (server *Server) initCancelReservationHandler(service *application.ReservationService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCancelReservationCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) startGrpcServer(reservationHandler *api.ReservationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reservation_service.RegisterReservationServiceServer(grpcServer, reservationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
