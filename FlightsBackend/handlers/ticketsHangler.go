package handlers

import (
	"context"
	"log"
	"net/http"
	"xwsproj/model"
	"xwsproj/repositories"

	"github.com/gorilla/mux"
)

type KeyTicket struct{}

type TicketsHandler struct {
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *repositories.TicketRepo
}

func NewTicketsHandler(l *log.Logger, t *repositories.TicketRepo) *TicketsHandler {
	return &TicketsHandler{l, t}
}
func (t *TicketsHandler) GetAllTickets(rw http.ResponseWriter, h *http.Request) {
	tickets, err := t.repo.GetAllTickets()
	if err != nil {
		t.logger.Print("Database exception: ", err)
	}

	if tickets == nil {
		return
	}

	err = tickets.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		t.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (t *TicketsHandler) GetTicketById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	ticket, err := t.repo.GetTicketById(id)
	if err != nil {
		t.logger.Print("Database exception: ", err)
	}

	if ticket == nil {
		http.Error(rw, "Ticket with given id not found", http.StatusNotFound)
		t.logger.Printf("Ticket with id: '%s' not found", id)
		return
	}

	err = ticket.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		t.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (t *TicketsHandler) GetTicketByUserId(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	ticket, err := t.repo.GetTicketByUserId(id)
	if err != nil {
		t.logger.Print("Database exception: ", err)
	}

	if ticket == nil {
		http.Error(rw, "Ticket with given user id not found", http.StatusNotFound)
		t.logger.Printf("Ticket with user id: '%s' not found", id)
		return
	}

	err = ticket.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		t.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (t *TicketsHandler) CreateTicket(rw http.ResponseWriter, h *http.Request) {
	ticket := h.Context().Value(KeyProduct{}).(*model.Ticket)
	t.repo.CreateTicket(ticket)
	rw.WriteHeader(http.StatusCreated)
}
func (t *TicketsHandler) DeleteTicket(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	t.repo.DeleteTicket(id)
	rw.WriteHeader(http.StatusNoContent)
}
func (t *TicketsHandler) PatchTicket(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	ticket := h.Context().Value(KeyProduct{}).(*model.Ticket)

	t.repo.Update(id, ticket)
	rw.WriteHeader(http.StatusOK)
}
func (t *TicketsHandler) MiddlewareTicketDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		ticket := &model.Ticket{}
		err := ticket.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			t.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, ticket)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
func (t *TicketsHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		t.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
