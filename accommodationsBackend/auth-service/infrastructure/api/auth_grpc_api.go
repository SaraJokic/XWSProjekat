package api

import (
	"accommodationsBackend/auth-service/application"
	"accommodationsBackend/auth-service/domain"
	"accommodationsBackend/auth-service/jwt"
	auth_service "accommodationsBackend/common/proto/auth-service"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
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
	/*fmt.Println("usao sam u insert metodu a ovo je ctx", ctx)

	md, err := handler.ReadMetadataFromContext(ctx)
	if err != nil {
		fmt.Println("Error reading metadata:", err)
		// Odlučite da li želite da vratite grešku ili da nastavite sa obradom zahteva
	} else {
		// Reading only the value associated with the "authorization" key from metadata
		if values, ok := md["authorization"]; ok && len(values) > 0 {
			// Typically, the Authorization header looks like "Bearer [token_value]".
			// Splitting it to get the actual token value.
			authParts := strings.Split(values[0], " ")
			if len(authParts) == 2 && strings.EqualFold(authParts[0], "Bearer") {
				token := authParts[1]
				fmt.Printf("OVO JE TOKEN IZ METAPODATAKA - %s\n", token)
			} else {
				fmt.Println("Invalid Authorization format")
				// Decide if you want to continue or return an error due to invalid format
			}
		} else {
			fmt.Println("Authorization token not found in metadata")
			// Decide if you want to continue or return an error due to missing token
		}
	}*/
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

func (handler *AuthHandler) ReadMetadataFromContext(ctx context.Context) (metadata.MD, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata found in context")
	}
	return md, nil
}

/*
func (handler *AuthHandler) Login(ctx context.Context, request *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {
	var user *domain.User
	user, err := handler.service.ValidateUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to validate username and password")
	}

	var claims = &jwt.JwtAccessTokenClaims{}
	claims.Id = user.Id
	claims.Name = user.Username
	claims.Username = request.Username
	claims.Role = user.Role

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)
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
//return &auth_service.LoginResponse{Token: tokenString}, nil

//}*/

func (handler *AuthHandler) Login(ctx context.Context, request *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {

	// Provera korisnickih podataka
	fetchedUser, userErr := handler.service.ValidateUsernameAndPassword(request.Username, request.Password)
	if userErr != nil {
		return nil, fmt.Errorf("error authenticating user")
	}

	// Kreiranje JWT tvrdnji na osnovu podataka iz fetchedUser
	tokenClaims := &jwt.JwtClaims{
		Id:       fetchedUser.Id,
		Name:     fetchedUser.Username,
		Username: request.Username,
		Role:     fetchedUser.Role,
	}

	// Generisnaje JWT tokena
	generatedToken, tokenErr := jwt.GenerateToken(tokenClaims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	// Token se vraca u odgovoru
	return &auth_service.LoginResponse{Token: generatedToken}, nil
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
func (handler *AuthHandler) GetAuthByUsername(ctx context.Context, request *auth_service.GetAuthRequest) (*auth_service.GetAuthResponse, error) {
	fmt.Println("EVO USAU U AUTH")
	user, err := handler.service.GetByUsername(request.Id)
	if err != nil {
		return nil, err
	}
	fmt.Println("AUTH: nadjen user: ", user)
	userMapped := mapUser(user)
	fmt.Println("AUTH: mapiran user:", userMapped)
	response := &auth_service.GetAuthResponse{
		User: userMapped,
	}
	return response, nil
}
func (handler *AuthHandler) DeleteAuthUser(ctx context.Context, request *auth_service.DeleteAuthRequest) (*auth_service.DeleteAuthResponse, error) {
	id := request.Id

	err := handler.service.Delete(id)
	if err != nil {
		return &auth_service.DeleteAuthResponse{Message: "User delete failed"}, nil
	}
	return &auth_service.DeleteAuthResponse{Message: "Auth User deleted"}, nil
}
