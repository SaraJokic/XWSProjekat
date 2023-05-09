package startup

import (
	"accommodationsBackend/accommodations-service/application"
	"accommodationsBackend/accommodations-service/domain"
	"accommodationsBackend/accommodations-service/infrastructure/api"
	"accommodationsBackend/accommodations-service/infrastructure/persistence"
	"accommodationsBackend/accommodations-service/startup/config"
	"accommodationsBackend/common/proto/accommodation_service"
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
	productStore := server.initAccommodationsStore(mongoClient)

	productService := server.initAccommodationService(productStore)

	productHandler := server.initAccommodationHandler(productService)

	server.startGrpcServer(productHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AccommodationDBHost, server.config.AccommodationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAccommodationsStore(client *mongo.Client) domain.AccommodationStore {
	store := persistence.NewAccommodationMongoDBStore(client)
	store.DeleteAll()
	for _, accommodation := range accommodations {
		err := store.Insert(accommodation)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initAccommodationService(store domain.AccommodationStore) *application.AccommodationService {
	return application.NewAccommodationService(store)
}

func (server *Server) initAccommodationHandler(service *application.AccommodationService) *api.AccommodationHandler {
	return api.NewAccommodationHandler(service)
}

func (server *Server) startGrpcServer(accHandler *api.AccommodationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accommodation_service.RegisterAccommodationServiceServer(grpcServer, accHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
