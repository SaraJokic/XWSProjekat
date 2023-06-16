package application

import (
	"accommodationsBackend/reservation-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	store domain.ReservationStore
}

func NewreservationService(store domain.ReservationStore) *ReservationService {
	return &ReservationService{
		store: store,
	}
}
func (service *ReservationService) Get(id primitive.ObjectID) (*domain.Reservation, error) {
	return service.store.Get(id)
}

func (service *ReservationService) GetAll() ([]*domain.Reservation, error) {
	return service.store.GetAll()
}
func (service *ReservationService) Update(id primitive.ObjectID, reservation *domain.Reservation) error {
	return service.store.Update(id, reservation)
}
func (service *ReservationService) Insert(reservation *domain.Reservation) error {
	if reservation.Id.IsZero() {
		reservation.Id = primitive.NewObjectID()
	}
	return service.store.Insert(reservation)
}
func (service *ReservationService) GetByAccommodationId(id primitive.ObjectID) ([]*domain.Reservation, error) {
	return service.store.GetByAccommodationId(id)

}
func (service *ReservationService) GetByUserId(id primitive.ObjectID) ([]*domain.Reservation, error) {
	return service.store.GetByUserId(id)

}
func (service *ReservationService) Delete(id string) error {
	return service.store.Delete(id)
}
