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

	// NoSQL: Initialize User Repository store
	userRepo, err := repositories.NewUserRepo(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer userRepo.DisconnectUserRepo(timeoutContext)

	// NoSQL: Checking if the connection was established
	flightRepo.PingFlightRepo()
	// NoSQL: Checking if the connection was established
	userRepo.PingUserRepo()

	//Initialize the handler and inject said logger
	flightsHandler := handlers.NewFlightsHandler(logger, flightRepo)
	//Initialize the handler and inject said logger
	usersHandler := handlers.NewUsersHandler(logger, userRepo)

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	//router.Use(flightsHandler.MiddlewareContentTypeSet)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", flightsHandler.GetAllFlights)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", flightsHandler.CreateFlight)
	postRouter.Use(flightsHandler.MiddlewareFlightDeserialization)

	getByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getByIdRouter.HandleFunc("/{id}", flightsHandler.GetFlightById)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id}", flightsHandler.DeleteFlight)

	postUserRouter := router.Methods(http.MethodPost).Subrouter()
	postUserRouter.HandleFunc("/register", usersHandler.PostUser)
	postUserRouter.Use(usersHandler.MiddlewareUserDeserialization)

	loginRouter := router.Methods(http.MethodPost).Subrouter()
	loginRouter.HandleFunc("/login", usersHandler.LoginUser)
	//loginRouter.Use(usersHandler.MiddlewareUserDeserialization)

	getAllUsersRouter := router.Methods(http.MethodGet).Subrouter()
	getAllUsersRouter.HandleFunc("/allUsers", usersHandler.GetAllUsers)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}))

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
