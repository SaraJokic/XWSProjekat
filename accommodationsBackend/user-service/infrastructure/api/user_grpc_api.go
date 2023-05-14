package api

import (
	"accommodationsBackend/common/proto/user_service"
	"accommodationsBackend/user-service/application"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	fmt.Println("usao sam u register funjciju, a oo je newUser", request.User.Name)
	exists, err := handler.service.CheckIfEmailAndUsernameExist(newUser.Email, newUser.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		response := &user_service.RegisterResponse{Message: "Email or username already exist"}
		return response, nil
	}

	/*	hashedPassword, err := p.HashPassword(newUser.Password)
		if err != nil {
			p.logger.Print("Error while hashing password:", err)
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		newUser.Password = hashedPassword*/

	hashedPassword, err := handler.HashPassword(newUser.Password)
	if err != nil {
		response := &user_service.RegisterResponse{Message: "Hashing password unsuccesfull"}
		return response, nil
	}
	newUser.Password = hashedPassword
	fmt.Println("ovo je newUser", newUser)

	userMapped := reverseMapUser(newUser)
	userMapped.Id = primitive.NewObjectID()
	fmt.Println("ovo je userMaped", userMapped)
	err = handler.service.Register(userMapped)
	if err != nil {
		return nil, err
	}
	/*	authUser := &domain.UserAuth{
			Username: newUser.Username,
			Password: newUser.Password,
		}
		authClient := servces.NewAuthClient("http://auth-service:8000/")
		mappedAuthUser := mapAuthUser(authUser)
		_, err = authClient.Insert(ctx, &auth_service.InsertRequest{Username: mappedAuthUser.Username, Password: mappedAuthUser.Password})
		if err != nil {
			return nil, err
		}

	*/
	response := &user_service.RegisterResponse{Message: "Registration successful!"}
	return response, nil

}
func (handler *UserHandler) HashPassword(password string) (string, error) {
	passwordbytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordbytes), err
}

func (handler *UserHandler) UpdateUser(ctx context.Context, request *user_service.UpdateRequest) (*user_service.UpdateResponse, error) {
	id := request.UserId // ID korisnika koji se ažurira
	user := request.User // Novi podaci korisnika
	fmt.Println("ovo je user sa id-em", id, user.Username)
	userMapped := reverseMapUser(user)
	fmt.Println("ovo je mapirani user", userMapped.Username)
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
