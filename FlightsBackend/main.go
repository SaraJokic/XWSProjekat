package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"xwsproj/handlers"
	"xwsproj/repositories"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//Reading from environment, if not set we will default it to 8080.
	//This allows flexibility in different environments (for eg. when running multiple docker api's and want to override the default port)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4200"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[flight-store] ", log.LstdFlags)

	// NoSQL: Initialize Product Repository store
	flightRepo, err := repositories.NewFlightRepo(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer flightRepo.DisconnectFlightRepo(timeoutContext)

	// NoSQL: Checking if the connection was established
	flightRepo.PingFlightRepo()

	//Initialize the handler and inject said logger
	flightsHandler := handlers.NewFlightsHandler(logger, flightRepo)

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	router.Use(flightsHandler.MiddlewareContentTypeSet)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", flightsHandler.GetAllFlights)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", flightsHandler.CreateFlight)
	postRouter.Use(flightsHandler.MiddlewareFlightDeserialization)

	getByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getByIdRouter.HandleFunc("/{id}", flightsHandler.GetFlightById)

	getByNameRouter := router.Methods(http.MethodGet).Subrouter()
	getByNameRouter.HandleFunc("/filter/{fromplace}", flightsHandler.GetFlightsFromPlace)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id}", flightsHandler.DeleteFlight)

	//cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))
	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
