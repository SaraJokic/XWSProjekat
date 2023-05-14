package application

import (
	"accommodationsBackend/accommodations-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	return service.store.Get(id)
}

func (service *AccommodationService) GetAll() ([]*domain.Accommodation, error) {
	return service.store.GetAll()
}

func (service *AccommodationService) Create(acc *domain.Accommodation) error {
	if acc.Id.IsZero() {
		acc.Id = primitive.NewObjectID()
	}
	return service.store.Insert(acc)
}
func (service *AccommodationService) DeleteAll() {
	service.store.DeleteAll()
}
