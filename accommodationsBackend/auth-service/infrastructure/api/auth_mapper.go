package api

import (
	"accommodationsBackend/auth-service/domain"
	auth_service "accommodationsBackend/common/proto/auth-service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapUser(user *domain.User) *auth_service.AuthUser {
	userMapped := &auth_service.AuthUser{
		Id:       user.Id.Hex(),
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}
	return userMapped
}

func reverseMapUser(user *auth_service.AuthUser) *domain.User {

	/*id, err := primitive.ObjectIDFromHex(user.Id)

	if err != nil {
		return nil
	}*/
	userMapped := &domain.User{

		Username: user.Username,
		Password: user.Password,
	}
	return userMapped
}
func reverseMapUserWithId(user *auth_service.AuthUser) *domain.User {

	id, err := primitive.ObjectIDFromHex(user.Id)

	if err != nil {
		return nil
	}
	userMapped := &domain.User{
		Id:       id,
		Username: user.Username,
		Password: user.Password,
	}
	return userMapped
}
