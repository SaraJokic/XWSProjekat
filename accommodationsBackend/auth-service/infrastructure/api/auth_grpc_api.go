package api

import (
	"accommodationsBackend/auth-service/application"
	"accommodationsBackend/auth-service/domain"
	"accommodationsBackend/auth-service/jwt"
	auth_service "accommodationsBackend/common/proto/auth-service"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AuthHandler struct {
	auth_service.UnimplementedAuthServiceServer
	service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (handler *AuthHandler) Insert(ctx context.Context, request *auth_service.InsertRequest) (*auth_service.InsertResponse, error) {
	newUser := &domain.User{
		//Id:       primitive.NewObjectID(),
		Username: request.Username,
		Password: request.Password,
		Role:     request.Role,
	}
	fmt.Println("ovo je newuseer u insert handleru", newUser.Username, newUser.Password)
	newUser.Id = primitive.NewObjectID()

	err := handler.service.Insert(newUser)
	if err != nil {
		return nil, err
	}
	response := &auth_service.InsertResponse{Message: "Registration successful!"}
	return response, nil

}

func (handler *AuthHandler) Login(ctx context.Context, request *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {

	var user *domain.User
	user, err := handler.service.ValidateUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to validate username and password")
	}

	var claims = &jwt.JwtClaims{}
	claims.Id = user.Id
	claims.Name = user.Username
	claims.Username = request.Username
	claims.Role = user.Role

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Second)
	tokenString, err := jwt.GenerateToken(claims, expirationTime)

	if err != nil {
		return nil, err
	}
	/*
		response := struct {
			Token string `json:"token"`
		}{
			Token: tokenString,
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}
	*/
	return &auth_service.LoginResponse{Token: tokenString}, nil

}
func (handler *AuthHandler) GetAll(ctx context.Context, request *auth_service.AllRequest) (*auth_service.AllResponse, error) {
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &auth_service.AllResponse{
		Users: []*auth_service.AuthUser{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}

	return response, nil
}

/*
func (handler *AuthHandler) ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Empty string", http.StatusUnauthorized)
			return
		}
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "split didnt work", http.StatusUnauthorized)
			return
		}
		tokenString := headerParts[1]

		valid, claims := jwt.VerifyToken(tokenString)
		if !valid {
			http.Error(w, "token not verified", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "Id", claims.Id)
		ctx = context.WithValue(ctx, "Name", claims.Name)
		ctx = context.WithValue(ctx, "Username", claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
*/
