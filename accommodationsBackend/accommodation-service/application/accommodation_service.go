package application

import (
	"accommodationsBackend/accommodations-service/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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

func (service *AccommodationService) GetAllProminentAccommodation() ([]*domain.Accommodation, error) {
	return service.store.GetAllProminentAccommodation()
}

func (service *AccommodationService) Create(acc *domain.Accommodation) error {
	if acc.Id.IsZero() {
		acc.Id = primitive.NewObjectID()
	}
	fmt.Println("U SERVISU SAM")
	return service.store.Insert(acc)
}
func (service *AccommodationService) DeleteAll() {
	service.store.DeleteAll()
}
func (service *AccommodationService) GetAccommodationByUserId(id primitive.ObjectID) ([]*domain.Accommodation, error) {
	return service.store.GetAccommodationByUserId(id)
}

func (service *AccommodationService) UpdateAccommodation(id string, accommodation *domain.Accommodation) error {
	return service.store.UpdateAccommodation(id, accommodation)
}
func (service *AccommodationService) GetAccommodationByLocation(location string) ([]*domain.Accommodation, error) {
	return service.store.GetAccommodationByLocation(location)
}
func (service *AccommodationService) GetAccommodationByNumberOfGuests(guest int) ([]*domain.Accommodation, error) {
	return service.store.GetAccommodationByNumberOfGuests(guest)
}
func (service *AccommodationService) GetAccommodationByDateRange(startDate, endDate time.Time) ([]*domain.Accommodation, error) {
	return service.store.GetAccommodationByDateRange(startDate, endDate)
}
