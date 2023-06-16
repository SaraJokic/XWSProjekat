package application

import (
	"accommodationsBackend/user-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}
func (service *UserService) Register(user *domain.User) error {
	return service.store.Register(user)
}

func (service *UserService) CheckIfEmailAndUsernameExist(email string, username string) (bool, error) {
	return service.store.CheckIfEmailAndUsernameExist(email, username)
}

func (service *UserService) UpdateUser(id string, user *domain.User) error {
	return service.store.UpdateUser(id, user)
}

func (service *UserService) Delete(id string) error {
	return service.store.Delete(id)
}
func (service *UserService) GetByUsername(username string) (*domain.User, error) {
	return service.store.GetByUsername(username)
}
