package startup

import (
	cfg "accommodationsBackend/api_gateway/startup/config"
	"accommodationsBackend/common/proto/accommodation_service"
	auth_service "accommodationsBackend/common/proto/auth-service"
	availability_service "accommodationsBackend/common/proto/availability-service"
	userGw "accommodationsBackend/common/proto/user_service"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	userEmdpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEmdpoint, opts)
	if err != nil {
		panic(err)
	}

	accommodationEmdpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err = accommodation_service.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEmdpoint, opts)
	if err != nil {
		panic(err)
	}

	availabilityEmdpoint := fmt.Sprintf("%s:%s", server.config.AvailabilityHost, server.config.AvailabilityPort)
	err = availability_service.RegisterAvailabilityServiceHandlerFromEndpoint(context.TODO(), server.mux, availabilityEmdpoint, opts)
	if err != nil {
		panic(err)
	}

	authEmdpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	err = auth_service.RegisterAuthServiceHandlerFromEndpoint(context.TODO(), server.mux, authEmdpoint, opts)
	if err != nil {
		panic(err)
	}
}

/*func (server *Server) initCustomHandlers() {
	catalogueEmdpoint := fmt.Sprintf("%s:%s", server.config.CatalogueHost, server.config.CataloguePort)
	orderingEmdpoint := fmt.Sprintf("%s:%s", server.config.OrderingHost, server.config.OrderingPort)
	shippingEmdpoint := fmt.Sprintf("%s:%s", server.config.ShippingHost, server.config.ShippingPort)
	orderingHandler := api.NewOrderingHandler(orderingEmdpoint, catalogueEmdpoint, shippingEmdpoint)
	orderingHandler.Init(server.mux)
}*/

func (server *Server) Start() {
	mainHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		server.mux.ServeHTTP(w, r)
	})

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(mainHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), corsHandler))

	//mux := http.NewServeMux()
	//mux.Handle("/", corsHandler)

	//handler := middleware.ValidateToken(mux)
}
