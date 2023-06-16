package api

import (
	auth_service "accommodationsBackend/common/proto/auth-service"
	"accommodationsBackend/common/proto/reservation_service"
	"accommodationsBackend/common/proto/user_service"
	"accommodationsBackend/user-service/application"
	"accommodationsBackend/user-service/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
)

type UserHandler struct {
	user_service.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *user_service.GetRequest) (*user_service.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	userMapped := mapUser(user)
	response := &user_service.GetResponse{
		User: userMapped,
	}
	return response, nil
}

func (handler *UserHandler) Register(ctx context.Context, request *user_service.RegisterRequest) (*user_service.RegisterResponse, error) {
	newUser := request.User
	exists, err := handler.service.CheckIfEmailAndUsernameExist(newUser.Email, newUser.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		response := &user_service.RegisterResponse{Message: "Email or username already exist"}
		return response, nil
	}

	hashedPassword, err := handler.HashPassword(newUser.Password)
	if err != nil {
		response := &user_service.RegisterResponse{Message: "Hashing password unsuccesfull"}
		return response, nil
	}
	newUser.Password = hashedPassword

	userMapped := reverseMapUser(newUser)
	userMapped.Id = primitive.NewObjectID()
	err = handler.service.Register(userMapped)
	if err != nil {
		return nil, err
	}
	client := NewAuthClient()

	client.Insert(context.Background(), &auth_service.InsertRequest{Username: userMapped.Username, Password: userMapped.Password, Role: handler.MapRole(userMapped.Role)})

	response := &user_service.RegisterResponse{Message: "Registration successful!"}
	return response, nil

}
func (handler *UserHandler) MapRole(role domain.UserType) string {
	if role == domain.Customer {
		return "Customer"
	}
	return "Admin"
}

func (handler *UserHandler) HashPassword(password string) (string, error) {
	passwordbytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordbytes), err
}
func NewAuthClient() auth_service.AuthServiceClient {
	conn, err := grpc.Dial("auth-service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return auth_service.NewAuthServiceClient(conn)
}
func (handler *UserHandler) UpdateUser(ctx context.Context, request *user_service.UpdateRequest) (*user_service.UpdateResponse, error) {
	id := request.UserId // ID korisnika koji se ažurira
	user := request.User // Novi podaci korisnika
	userMapped := reverseMapUser(user)
	err := handler.service.UpdateUser(id, userMapped) // Ažuriranje korisnika u repozitorijumu
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update user") // Vraćanje greške ako ažuriranje nije uspelo
	}

	return &user_service.UpdateResponse{Message: "Update succesfull"}, nil // Vraćanje potvrde da je ažuriranje uspelo
}

/*
	func (handler *UserHandler) DeleteUser(ctx context.Context, request *user_service.DeleteRequest) (*user_service.DeleteResponse, error) {
		id := request.UserId
		fmt.Println("ovo je userid", id)
		err := handler.service.Delete(id)
		if err != nil {
			return &user_service.DeleteResponse{Message: "User delete failed"}, nil
		}

		return &user_service.DeleteResponse{Message: "User deleted"}, nil
	}
*/
func (handler *UserHandler) DeleteUser(ctx context.Context, request *user_service.DeleteRequest) (*user_service.DeleteResponse, error) {
	id := request.Id

	user, _ := handler.Get(context.Background(), &user_service.GetRequest{Id: id})
	fmt.Println("USER: ovo je nadjen user:", user)
	fmt.Println("Username nadjenog usera: ", user.User.Username)
	client := NewAuthClient()
	auth, _ := client.GetAuthByUsername(context.Background(), &auth_service.GetAuthRequest{Id: user.User.Username})
	fmt.Println("USER: nadjen auth user: ", auth)
	client.DeleteAuthUser(context.Background(), &auth_service.DeleteAuthRequest{Id: auth.User.Id})

	rClient := NewReservationClient()
	reservations, _ := rClient.GetReservationByUserId(context.Background(), &reservation_service.GetReservationByUserIdRequest{
		Id: user.User.Id,
	})
	for _, r := range reservations.Reservations {
		_, err := rClient.DeleteReservation(context.Background(), &reservation_service.DeleteReservationRequest{Id: r.Id})
		if err != nil {
			// Handle the error
			fmt.Printf("Failed to delete reservation with ID %s: %v\n", r.Id, err)
		}
	}

	err := handler.service.Delete(id)
	if err != nil {
		return &user_service.DeleteResponse{Message: "User delete failed"}, nil
	}
	return &user_service.DeleteResponse{Message: "User deleted"}, nil
}
func (handler *UserHandler) GetAll(ctx context.Context, request *user_service.GetAllRequest) (*user_service.GetAllResponse, error) {
	users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &user_service.GetAllResponse{
		Users: []*user_service.User{},
	}
	for _, user := range users {
		current := mapUser(user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
func (handler *UserHandler) GetByUsername(ctx context.Context, request *user_service.GetRequest) (*user_service.GetResponse, error) {
	user, err := handler.service.GetByUsername(request.Id)
	if err != nil {
		return nil, err
	}
	userMapped := mapUser(user)
	response := &user_service.GetResponse{
		User: userMapped,
	}
	return response, nil
}
func NewReservationClient() reservation_service.ReservationServiceClient {
	conn, err := grpc.Dial("reservation-service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return reservation_service.NewReservationServiceClient(conn)
}
