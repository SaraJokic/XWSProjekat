package api

import (
	"accommodationsBackend/common/proto/user_service"
	"accommodationsBackend/user-service/domain"
)

func mapUser(user *domain.User) *user_service.User {
	userMapped := &user_service.User{
		Id:      user.Id.Hex(),
		Name:    user.Name,
		Surname: user.Surname,
	}
	return userMapped
}
