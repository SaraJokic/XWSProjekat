package startup

import (
	"accommodationsBackend/common/proto/user_service"
	saga "accommodationsBackend/common/saga/messaging"
	"accommodationsBackend/common/saga/messaging/nats"
	"accommodationsBackend/user-service/application"
	"accommodationsBackend/user-service/domain"
	"accommodationsBackend/user-service/infrastructure/api"
	"accommodationsBackend/user-service/infrastructure/persistence"
	"accommodationsBackend/user-service/middleware"
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

const (
	QueueGroup = "user_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	usersStore := server.initUsersStore(mongoClient)

	userService := server.initUserService(usersStore)

	commandSubscriber := server.initSubscriber(server.config.CancelReservationCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CancelReservationReplySubject)
	server.initCancelReservationHandler(userService, replyPublisher, commandSubscriber)

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

	return store
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
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

func (server *Server) initCancelReservationHandler(service *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCancelReservationCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	interceptor := middleware.NewAuthInterceptor()

	return api.NewUserHandler(service, interceptor)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	interceptor := middleware.NewAuthInterceptor()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))
	user_service.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
