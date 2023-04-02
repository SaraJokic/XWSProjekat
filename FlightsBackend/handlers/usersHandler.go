package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"xwsproj/jwt"
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

func (p *UsersHandler) DeleteUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	p.repo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}

func (p *UsersHandler) PostUser(rw http.ResponseWriter, h *http.Request) {
	newUser := h.Context().Value(KeyUser{}).(*model.User)

	exists, err := p.repo.CheckIfEmailAndUsernameExist(newUser.Email, newUser.Username)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(rw, "Email or username already exist", http.StatusBadRequest)
		return
	}

	hashedPassword, err := p.HashPassword(newUser.Password)
	if err != nil {
		p.logger.Print("Error while hashing password:", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	newUser.Password = hashedPassword
	err = p.repo.Insert(newUser)
	if err != nil {
		p.logger.Print("Error while inserting user:", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (p *UsersHandler) HashPassword(password string) (string, error) {
	passwordbytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordbytes), err
}

func (p *UsersHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginObj model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginObj)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	// validate the loginObj for valid credential adn if these are valid then
	var user *model.User
	user, err = p.repo.ValidateUsernameAndPassword(loginObj.Username, loginObj.Password)
	if err != nil {
		http.Error(w, "Username or password are incorrect", http.StatusBadRequest)
		return
	}

	var claims = &jwt.JwtClaims{}
	claims.Id = user.ID
	claims.Username = loginObj.Username
	claims.Role = user.Role
	claims.Name = user.Name

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Second)
	tokenString, err := jwt.GenerateToken(claims, expirationTime)

	if err != nil {
		http.Error(w, "error in generating token: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))

}

func (p *UsersHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		p.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
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
