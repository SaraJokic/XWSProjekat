package startup

import (
	"accommodationsBackend/accommodations-service/domain"
	"accommodationsBackend/accommodations-service/infrastructure/api"
	"accommodationsBackend/accommodations-service/infrastructure/persistence"
	"accommodationsBackend/accommodations-service/startup/config"

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
	productStore := server.initUsersStore(mongoClient)

	productService := server.initUserService(productStore)

	productHandler := server.initUserHandler(productService)

	server.startGrpcServer(productHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUsersStore(client *mongo.Client) domain.AccommodationStore {
	store := persistence.NewUserMongoDBStore(client)
	store.DeleteAll()
	for _, user := range accommodations {
		err := store.Insert(user)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initUserService(store domain.AccommodationStore) *application.AccommodationService {
	return application.NewUserService(store)
}

func (server *Server) initUserHandler(service *application.AccommodationService) *api.UserHandler {
	return api.NewProductHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user_service.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
