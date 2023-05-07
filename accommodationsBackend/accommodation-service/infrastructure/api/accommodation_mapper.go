package api

import (
	"accommodationsBackend/accommodations-service/domain"
)

func mapAccommodation(acc *domain.Accommodation) *accommodation_service.Accommodation {
	accMapped := &accommodation_service.Accommodation{
		Id:        acc.Id.Hex(),
		Name:      acc.Name,
		Location:  acc.Location,
		Benefits:  acc.Benefits,
		MinGuests: acc.MinGuests,
		MaxGuests: acc.MaxGuests,
	}
	return accMapped
}
