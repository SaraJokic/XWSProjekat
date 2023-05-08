package api

import (
	"accommodationsBackend/accommodations-service/domain"
	"accommodationsBackend/common/proto/accommodation_service"
)

func mapBenefits(benefits domain.Benefits) *accommodation_service.Benefits {
	benefitsMapped := &accommodation_service.Benefits{
		Id:          benefits.ID.String(),
		Wifi:        benefits.Wifi,
		Kitchen:     benefits.Kitchen,
		FreeParking: benefits.FreeParking,
	}
	return benefitsMapped
}

func mapAccommodation(acc *domain.Accommodation) *accommodation_service.Accommodation {
	accMapped := &accommodation_service.Accommodation{
		Id:        acc.Id.Hex(),
		Name:      acc.Name,
		Location:  acc.Location,
		Benefits:  mapBenefits(acc.Benefits),
		MinGuests: int32(acc.MinGuests),
		MaxGuests: int32(acc.MaxGuests),
	}
	return accMapped
}
