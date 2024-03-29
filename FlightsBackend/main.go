package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"xwsproj/handlers"
	"xwsproj/middleware"
	"xwsproj/repositories"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//Reading from environment, if not set we will default it to 8080.
	//This allows flexibility in different environments (for eg. when running multiple docker api's and want to override the default port)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
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

	// NoSQL: Initialize User Repository store
	userRepo, err := repositories.NewUserRepo(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer userRepo.DisconnectUserRepo(timeoutContext)

	ticketRepo, err := repositories.NewTicketRepo(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer ticketRepo.DisconnectTicketRepo(timeoutContext)

	// NoSQL: Checking if the connection was established
	flightRepo.PingFlightRepo()
	// NoSQL: Checking if the connection was established
	userRepo.PingUserRepo()
	ticketRepo.PingTicketRepo()

	//Initialize the handler and inject said logger
	flightsHandler := handlers.NewFlightsHandler(logger, flightRepo)
	//Initialize the handler and inject said logger
	usersHandler := handlers.NewUsersHandler(logger, userRepo)
	ticketsHandler := handlers.NewTicketsHandler(logger, ticketRepo)

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	//router.Use(flightsHandler.MiddlewareContentTypeSet)

	//FLIGHT

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", flightsHandler.GetAllFlights)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", flightsHandler.CreateFlight)
	postRouter.Use(flightsHandler.MiddlewareFlightDeserialization)
	postRouter.Use(middleware.ValidateToken)

	getByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getByIdRouter.HandleFunc("/{id}", flightsHandler.GetFlightById)
	getByIdRouter.Use(middleware.ValidateToken)

	patchRouter := router.Methods(http.MethodPatch).Subrouter()
	patchRouter.HandleFunc("/{id}", flightsHandler.UpdateFlight)
	patchRouter.Use(flightsHandler.MiddlewareFlightDeserialization)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id}", flightsHandler.DeleteFlight)
	deleteRouter.Use(middleware.ValidateToken)

	//SEARCHING FLIGHTS
	getByNameRouter := router.Methods(http.MethodGet).Subrouter()
	getByNameRouter.HandleFunc("/flight/filter", flightsHandler.GetFlightsFromPlaceToPlace)

	getBySeatsRouter := router.Methods(http.MethodGet).Subrouter()
	getBySeatsRouter.HandleFunc("/flight/filter/seats", flightsHandler.GetFlightsByNumOfSeats)

	/*
		getByStartDateRouter := router.Methods(http.MethodGet).Subrouter()
		getByStartDateRouter.HandleFunc("/flight/filter/time", flightsHandler.GetFlightsByStartTime)

		getSearchedFlightsRouter := router.Methods(http.MethodPatch).Subrouter()
		getSearchedFlightsRouter.HandleFunc("/flight/search", flightsHandler.GetSearchedFlights)
		getSearchedFlightsRouter.Use(flightsHandler.MiddlewareFlightSearchDeserialization)

	*/

	//USER

	postUserRouter := router.Methods(http.MethodPost).Subrouter()
	postUserRouter.HandleFunc("/users/register", usersHandler.PostUser)
	postUserRouter.Use(usersHandler.MiddlewareUserDeserialization)

	loginRouter := router.Methods(http.MethodPost).Subrouter()
	loginRouter.HandleFunc("/login", usersHandler.LoginUser)

	getAllUsersRouter := router.Methods(http.MethodGet).Subrouter()
	getAllUsersRouter.HandleFunc("/users/getAll", usersHandler.GetAllUsers)
	getAllUsersRouter.Use(middleware.ValidateToken)

	ValidateTokenRouter := router.Methods(http.MethodGet).Subrouter()
	ValidateTokenRouter.HandleFunc("/users/token/valildation", usersHandler.GetAllUsers)
	ValidateTokenRouter.Use(middleware.ValidateToken)

	//TICKETS

	getAllTicketsRouter := router.Methods(http.MethodGet).Subrouter()
	getAllTicketsRouter.HandleFunc("/tickets/all", ticketsHandler.GetAllTickets)
	getAllTicketsRouter.Use(middleware.ValidateToken)

	getTicketByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getTicketByIdRouter.HandleFunc("/tickets/get/{id}", ticketsHandler.GetTicketById)
	getTicketByIdRouter.Use(middleware.ValidateToken)

	getTicketByUserIdRouter := router.Methods(http.MethodGet).Subrouter()
	getTicketByUserIdRouter.HandleFunc("/tickets/getbyuser/{id}", ticketsHandler.GetTicketByUserId)
	getTicketByUserIdRouter.Use(middleware.ValidateToken)

	postTicketRouter := router.Methods(http.MethodPost).Subrouter()
	postTicketRouter.HandleFunc("/tickets/buy", ticketsHandler.CreateTicket)
	postTicketRouter.Use(ticketsHandler.MiddlewareTicketDeserialization)

	patchTicketRouter := router.Methods(http.MethodPatch).Subrouter()
	patchTicketRouter.HandleFunc("/tickets/update/{id}", ticketsHandler.PatchTicket)
	patchTicketRouter.Use(ticketsHandler.MiddlewareTicketDeserialization)
	postTicketRouter.Use(middleware.ValidateToken)

	deleteTicketRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteTicketRouter.HandleFunc("/tickets/delete/{id}", ticketsHandler.DeleteTicket)
	deleteRouter.Use(middleware.ValidateToken)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
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
