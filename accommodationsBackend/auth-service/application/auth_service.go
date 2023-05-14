package application

import (
	"accommodationsBackend/auth-service/domain"
)

type AuthService struct {
	store domain.AuthStore
}

func NewAuthService(store domain.AuthStore) *AuthService {
	return &AuthService{
		store: store,
	}

}

func (service *AuthService) Insert(user *domain.User) error {
	return service.store.Insert(user)
}

func (service *AuthService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}
func (service *AuthService) ValidateUsernameAndPassword(username string, password string) (*domain.User, error) {
	return service.store.ValidateUsernameAndPassword(username, password)
}
