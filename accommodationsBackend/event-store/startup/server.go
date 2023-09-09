package startup

import (
	"accommodationsBackend/common/proto/eventstore"
	nats "accommodationsBackend/common/saga/messaging/nats"
	"accommodationsBackend/event-store/domain"
	"accommodationsBackend/event-store/infrastructure/api"
	"accommodationsBackend/event-store/infrastructure/persistence"
	"accommodationsBackend/event-store/startup/config"
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
func (server *Server) Start() {
	natsComp := nats.GetNATSComponent()
	fmt.Printf("Nats komponenta u Event Store Serveru: %v\n", natsComp)
	nats.CreateJStream(natsComp, "Accommodations", "Accommodations.*")
	mongoClient := server.initMongoClient()
	eventStore := server.initEventStore(mongoClient)
	//natsComp := nats.NewNATSComponent("eventstore-service")
	//natsComp.ConnectToServer("nats://nats:4222")
	handler := initEventHandler(eventStore, natsComp)

	server.startGrpcServer(handler)
}
func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.EventsDBHost, server.config.EventsDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func (server *Server) initEventStore(client *mongo.Client) domain.EventStore {
	store := persistence.NewEventsnMongoDBStore(client)
	return store
}
func initEventHandler(store domain.EventStore, nats *nats.NATSComponent) *api.EventStoreHandler {
	return api.NewEventStoreHandler(store, nats)
}

func (server *Server) startGrpcServer(eventsHandler *api.EventStoreHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	eventstore.RegisterEventStoreServer(grpcServer, eventsHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
