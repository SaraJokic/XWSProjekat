package api

import (
	"accommodationsBackend/common/proto/user_service"
	"accommodationsBackend/user-service/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapUser(user *domain.User) *user_service.User {
	userMapped := &user_service.User{
		Id:       user.Id.Hex(),
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Name:     user.Name,
		LastName: user.LastName,
		City:     user.City,
		Country:  user.Country,
		Role:     user_service.UserType(user.Role),
	}
	return userMapped
}
func reverseMapUser(user *user_service.User) *domain.User {

	fmt.Println("usao sam u reversemapper funjciju")
	fmt.Println("userid", user.Id)
	/*id, err := primitive.ObjectIDFromHex(user.Id)

	if err != nil {
		return nil
	}*/
	userMapped := &domain.User{

		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Name:     user.Name,
		LastName: user.LastName,
		City:     user.City,
		Country:  user.Country,
		Role:     domain.UserType(user.Role),
	}
	fmt.Println("ovo je user u reversemapper-u", userMapped)
	return userMapped
}
func reverseMapUserWithId(user *user_service.User) *domain.User {

	fmt.Println("usao sam u reversemapper funjciju")
	fmt.Println("userid", user.Id)
	id, err := primitive.ObjectIDFromHex(user.Id)

	if err != nil {
		return nil
	}
	userMapped := &domain.User{
		Id:       id,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Name:     user.Name,
		LastName: user.LastName,
		City:     user.City,
		Country:  user.Country,
		Role:     domain.UserType(user.Role),
	}
	fmt.Println("ovo je user u reversemapper-u", userMapped)
	return userMapped
}

/*
func mapAuthUser(user *domain.UserAuth) *user_service.UserAuth {
	userMapped := &user_service.UserAuth{
		Username: user.Username,
		Password: user.Password,
	}
	return userMapped
}
*/
