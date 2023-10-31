package startup

import (
	application "accommodationsBackend/availability-service/application"
	"accommodationsBackend/availability-service/domain"
	"accommodationsBackend/availability-service/infrastructure/api"
	"accommodationsBackend/availability-service/infrastructure/persistence"
	"accommodationsBackend/availability-service/startup/config"
	availability_service "accommodationsBackend/common/proto/availability-service"
	saga "accommodationsBackend/common/saga/messaging"
	"accommodationsBackend/common/saga/messaging/nats"
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
	QueueGroup = "availability_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	availabilitiesStore := server.initAvailabilityStore(mongoClient)

	availabilitiesService := server.initAvailabilityService(availabilitiesStore)
	natsComp := nats.GetNATSComponent()
	fmt.Printf("Nats komponenta u Availability Serveru: %v\n", natsComp)
	/*natsComp := nats.NewNATSComponent("availability-service")
	err := natsComp.ConnectToServer("nats://nats:4222")
	if err != nil {
		log.Println("Greska pri konektovanju na server.")
	} else {
		log.Println("Uspesno konektovan na server")
	}*/

	//commandSubscriber := server.initSubscriber(natsComp, server.config.CancelReservationCommandSubject, QueueGroup)
	//replyPublisher := server.initPublisher(natsComp, server.config.CancelReservationReplySubject)
	//server.initCancelReservationHandler(availabilitiesService, replyPublisher, commandSubscriber)
	server.initEventHandler(natsComp, availabilitiesService)
	availabilityHandler := server.initAvailabilityHandler(availabilitiesService)

	server.startGrpcServer(availabilityHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AvailabilityDBHost, server.config.AvailabilityDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAvailabilityStore(client *mongo.Client) domain.AvailabilityStore {
	store := persistence.NewAvailabilityMongoDBStore(client)
	store.DeleteAll()
	for _, a := range availabilities {
		err := store.Insert(a)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initAvailabilityService(store domain.AvailabilityStore) *application.AvailabilityService {
	return application.NewAvailabilityService(store)
}

func (server *Server) initPublisher(natsComp *nats.NATSComponent, subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		natsComp, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(natsComp *nats.NATSComponent, subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		natsComp, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initCancelReservationHandler(service *application.AvailabilityService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCancelReservationCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initEventHandler(nc *nats.NATSComponent, availabilityService *application.AvailabilityService) {
	_, err := api.NewEventHandler(nc, availabilityService)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initAvailabilityHandler(service *application.AvailabilityService) *api.AvailabilityHandler {
	return api.NewAvailabilityHandler(service)
}

func (server *Server) startGrpcServer(availabilityHandler *api.AvailabilityHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	availability_service.RegisterAvailabilityServiceServer(grpcServer, availabilityHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
