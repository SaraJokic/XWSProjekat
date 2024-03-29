package application

import (
	"accommodationsBackend/rating-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingsService struct {
	store domain.RatingsStore
}

func NewRatingService(store domain.RatingsStore) *RatingsService {
	return &RatingsService{
		store: store,
	}
}
func (service *RatingsService) GetRateHost(id primitive.ObjectID) (*domain.RateHost, error) {
	return service.store.GetRateHost(id)
}

func (service *RatingsService) GetRateAccommodation(id primitive.ObjectID) (*domain.RateAccommodation, error) {
	return service.store.GetRateAccommodation(id)
}

func (service *RatingsService) GetAccommodationsRatingsByGuestId(id primitive.ObjectID) ([]*domain.RateAccommodation, error) {
	return service.store.GetAccommodationsRatingsByGuestId(id)
}

func (service *RatingsService) GetHostRatingsByGuestId(id primitive.ObjectID) ([]*domain.RateHost, error) {
	return service.store.GetHostRatingsByGuestId(id)
}

func (service *RatingsService) GetAccommodationsRatingsByAccommodationId(id primitive.ObjectID) ([]*domain.RateAccommodation, error) {
	return service.store.GetAccommodationsRatingsByAccommodationId(id)
}

func (service *RatingsService) GetHostRatingsByHostId(id primitive.ObjectID) ([]*domain.RateHost, error) {
	return service.store.GetHostRatingsByHostId(id)
}

func (service *RatingsService) CreateRateHost(host *domain.RateHost) error {
	if host.Id.IsZero() {
		host.Id = primitive.NewObjectID()
	}
	return service.store.InsertHostRating(host)
}
func (service *RatingsService) CreateRateAccommodation(acc *domain.RateAccommodation) error {
	if acc.Id.IsZero() {
		acc.Id = primitive.NewObjectID()
	}
	return service.store.InsertAccommodationRating(acc)
}

func (service *RatingsService) UpdateAccommodationRating(id string, rating *domain.RateAccommodation) error {
	return service.store.UpdateAccommodationRating(id, rating)
}

func (service *RatingsService) UpdateHostRating(id string, rating *domain.RateHost) error {
	return service.store.UpdateHostRating(id, rating)
}

func (service *RatingsService) DeleteAll() {
	service.store.DeleteAll()
}
func (service *RatingsService) DeleteRateHost(id primitive.ObjectID) error {
	return service.store.DeleteRateHost(id)
}
func (service *RatingsService) DeleteRateAccommodation(id primitive.ObjectID) error {
	return service.store.DeleteRateAccommodation(id)
}
