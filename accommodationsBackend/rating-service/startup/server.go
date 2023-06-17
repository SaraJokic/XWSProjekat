package startup

import (
	"accommodationsBackend/common/proto/rating_service"
	"accommodationsBackend/rating-service/application"
	"accommodationsBackend/rating-service/domain"
	"accommodationsBackend/rating-service/infrastructure/api"
	"accommodationsBackend/rating-service/infrastructure/persistence"
	"accommodationsBackend/rating-service/startup/config"
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
	ratingsStore := server.initRatingsStore(mongoClient)

	ratingsService := server.initRatingsService(ratingsStore)

	ratingsHandler := server.initRatingsHandler(ratingsService)

	server.startGrpcServer(ratingsHandler)
}
func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.RatingsDBHost, server.config.RatingsDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func (server *Server) initRatingsStore(client *mongo.Client) domain.RatingsStore {
	store := persistence.NewRatingMongoDBStore(client)
	/*store.DeleteAll()
	for _, a := range availabilities {
		err := store.Insert(a)
		if err != nil {
			log.Fatal(err)
		}
	}*/
	return store
}
func (server *Server) initRatingsService(store domain.RatingsStore) *application.RatingsService {
	return application.NewRatingService(store)
}
func (server *Server) initRatingsHandler(service *application.RatingsService) *api.RatingsHandler {
	return api.NewRatingsHandler(service)
}
func (server *Server) startGrpcServer(handler *api.RatingsHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	rating_service.RegisterRatingServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
