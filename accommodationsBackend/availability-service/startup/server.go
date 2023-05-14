package startup

import (
	application "accommodationsBackend/availability-service/application"
	"accommodationsBackend/availability-service/domain"
	"accommodationsBackend/availability-service/infrastructure/api"
	"accommodationsBackend/availability-service/infrastructure/persistence"
	"accommodationsBackend/availability-service/startup/config"
	availability_service "accommodationsBackend/common/proto/availability-service"
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
	mongoClient := server.initMongoClient()
	availabilitiesStore := server.initAvailabilityStore(mongoClient)

	availabilitiesService := server.initAvailabilityService(availabilitiesStore)

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
