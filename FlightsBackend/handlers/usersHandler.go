package handlers

import (
	"context"
	"log"
	"net/http"
	"xwsproj/model"
	"xwsproj/repositories"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type KeyUser struct{}

type UsersHandler struct {
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *repositories.UserRepo
}

func NewUsersHandler(l *log.Logger, r *repositories.UserRepo) *UsersHandler {
	return &UsersHandler{l, r}
}
func (p *UsersHandler) GetAllUsers(rw http.ResponseWriter, h *http.Request) {
	users, err := p.repo.GetAll()
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *UsersHandler) GetUserById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	user, err := p.repo.GetById(id)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if user == nil {
		http.Error(rw, "User with given id not found", http.StatusNotFound)
		p.logger.Printf("User with id: '%s' not found", id)
		return
	}

	err = user.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}
func (p *UsersHandler) PostUser(rw http.ResponseWriter, h *http.Request) {
	user := h.Context().Value(KeyUser{}).(*model.User)
	user.Password, _ = p.HashPassword(user.Password)
	p.repo.Insert(user)
	rw.WriteHeader(http.StatusCreated)
}

func (p *UsersHandler) HashPassword(password string) (string, error) {
	passwordbytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordbytes), err
}
func (p *UsersHandler) DeleteUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	p.repo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}
func (p *UsersHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		user := &model.User{}
		err := user.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			p.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyUser{}, user)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
func (p *UsersHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		p.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
