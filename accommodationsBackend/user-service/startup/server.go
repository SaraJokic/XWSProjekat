package startup

import (
	"accommodationsBackend/common/proto/user_service"
	"accommodationsBackend/user-service/application"
	"accommodationsBackend/user-service/domain"
	"accommodationsBackend/user-service/infrastructure/api"
	"accommodationsBackend/user-service/infrastructure/persistence"
	"accommodationsBackend/user-service/startup/config"
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
	usersStore := server.initUsersStore(mongoClient)

	userService := server.initUserService(usersStore)

	userHandler := server.initUserHandler(userService)

	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUsersStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	store.DeleteAll()
	/*for _, user := range users {
		err := store.Insert(user)
		if err != nil {
			log.Fatal(err)
		}
	}*/
	return store
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
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
