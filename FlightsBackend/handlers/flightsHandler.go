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

/*
func (f *FlightsHandler) GetSearchedFlights(rw http.ResponseWriter, req *http.Request) {
	flightsearchDTO := req.Context().Value(KeyProduct{}).(*model.FlightSearchDTO)
	flights, err := f.repo.GetSearched(flightsearchDTO)
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}

	if flights == nil {
		return
	}

	err = flights.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

*/

func (f *FlightsHandler) MiddlewareFlightSearchDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		flight := &model.FlightSearchDTO{}
		err := flight.FromJSON(req.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			f.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(req.Context(), KeyProduct{}, flight)
		req = req.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}

/*
func (p *FlightsHandler) GetFlightsByStartTime(rw http.ResponseWriter, h *http.Request) {
	starttime := h.URL.Query().Get("starttime")

	flight, err := p.repo.GetFlightsByStartTime(starttime)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if flight == nil {
		http.Error(rw, "Flight with given fromPlace not found", http.StatusNotFound)
		p.logger.Printf("Flight with fromPlace: '%s' not found", starttime)
		return
	}

	err = flight.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

*/

func (p *FlightsHandler) GetFlightsByNumOfSeats(rw http.ResponseWriter, h *http.Request) {
	numofseats := h.URL.Query().Get("numofseats")

	//result, _ := strconv.ParseInt(numofseats, 10, 64)
	//result := strconv.Itoa(numofseats)

	flight, err := p.repo.GetFlightsByNumOfSeats(numofseats)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if flight == nil {
		http.Error(rw, "Flight with given fromPlace not found", http.StatusNotFound)
		p.logger.Printf("Flight with fromPlace: '%s' not found", numofseats)
		return
	}

	err = flight.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *FlightsHandler) GetFlightsFromPlaceToPlace(rw http.ResponseWriter, h *http.Request) {
	fromplace := h.URL.Query().Get("fromplace")
	toplace := h.URL.Query().Get("toplace")

	//flight, err := p.repo.GetFlightsFromPlace(fromplace)
	flight, err := p.repo.GetFlightsFromPlaceToPlace(fromplace, toplace)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if flight == nil {
		http.Error(rw, "Flight with given fromPlace not found", http.StatusNotFound)
		//p.logger.Printf("Flight with fromPlace: '%s' not found", fromplace)
		p.logger.Printf("Flight with fromPlace: '%s' to place '%s' not found", fromplace, toplace)
		return
	}

	err = flight.ToJSON(rw)
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

func (p *FlightsHandler) UpdateFlight(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	flight := h.Context().Value(KeyProduct{}).(*model.Flight)

	p.repo.UpdateFlight(id, flight)
	rw.WriteHeader(http.StatusOK)
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
