package handlers

import (
	"context"
	"log"
	"net/http"
	"xwsproj/model"
	"xwsproj/repositories"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type FlightsHandler struct {
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *repositories.FlightRepo
}

func NewFlightsHandler(l *log.Logger, r *repositories.FlightRepo) *FlightsHandler {
	return &FlightsHandler{l, r}
}
func (p *FlightsHandler) GetAllFlights(rw http.ResponseWriter, h *http.Request) {
	flights, err := p.repo.GetAllFlights()
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if flights == nil {
		return
	}

	err = flights.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *FlightsHandler) GetFlightById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	flight, err := p.repo.GetFlightById(id)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if flight == nil {
		http.Error(rw, "Flight with given id not found", http.StatusNotFound)
		p.logger.Printf("Flight with id: '%s' not found", id)
		return
	}

	err = flight.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *FlightsHandler) GetFlightsFromPlace(rw http.ResponseWriter, h *http.Request) {
	fromplace := h.URL.Query().Get("fromplace")

	flights, err := p.repo.GetFlightsFromPlace(fromplace)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if flights == nil {
		return
	}

	err = flights.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *FlightsHandler) CreateFlight(rw http.ResponseWriter, h *http.Request) {
	flight := h.Context().Value(KeyProduct{}).(*model.Flight)
	p.repo.CreateFlight(flight)
	rw.WriteHeader(http.StatusCreated)
}
func (p *FlightsHandler) DeleteFlight(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	p.repo.DeleteFlight(id)
	rw.WriteHeader(http.StatusNoContent)
}
func (p *FlightsHandler) MiddlewareFlightDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		flight := &model.Flight{}
		err := flight.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			p.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, flight)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
func (p *FlightsHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		p.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
