package startup

import (
	"accommodationsBackend/common/proto/reservation_service"
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

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	reservationsStore := server.initReservationStore(mongoClient)

	reservationsService := server.initReservationService(reservationsStore)

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

func (server *Server) initReservationService(store domain.ReservationStore) *application.ReservationService {
	return application.NewreservationService(store)
}

func (server *Server) initReservationHandler(service *application.ReservationService) *api.ReservationHandler {
	return api.NewReservationHandler(service)
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
