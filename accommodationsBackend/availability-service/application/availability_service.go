package application

import (
	domain "accommodationsBackend/availability-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvailabilityService struct {
	store domain.AvailabilityStore
}

func NewAvailabilityService(store domain.AvailabilityStore) *AvailabilityService {
	return &AvailabilityService{
		store: store,
	}
}

func (service *AvailabilityService) Get(id primitive.ObjectID) (*domain.Availability, error) {
	return service.store.Get(id)
}

func (service *AvailabilityService) GetAll() ([]*domain.Availability, error) {
	return service.store.GetAll()
}
func (service *AvailabilityService) Update(id primitive.ObjectID, availability *domain.Availability) error {
	return service.store.Update(id, availability)
}
func (service *AvailabilityService) GetByAccommodationId(id primitive.ObjectID) (*domain.Availability, error) {
	return service.store.GetByAccommodationId(id)
}
func (service *AvailabilityService) Insert(availability *domain.Availability) error {
	if availability.Id.IsZero() {
		availability.Id = primitive.NewObjectID()
	}
	return service.store.Insert(availability)
}
