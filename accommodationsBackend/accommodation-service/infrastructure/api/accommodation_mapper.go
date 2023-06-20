package api

import (
	"accommodationsBackend/accommodations-service/domain"
	"accommodationsBackend/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapBenefits(benefits domain.Benefits) *accommodation_service.Benefits {
	benefitsMapped := &accommodation_service.Benefits{
		Wifi:        benefits.Wifi,
		Kitchen:     benefits.Kitchen,
		FreeParking: benefits.FreeParking,
	}
	return benefitsMapped
}

func mapAccommodation(acc *domain.Accommodation) *accommodation_service.Accommodation {
	accMapped := &accommodation_service.Accommodation{
		Id:               acc.Id.Hex(),
		HostId:           acc.HostId.Hex(),
		Name:             acc.Name,
		Location:         acc.Location,
		Benefits:         mapBenefits(acc.Benefits),
		MinGuests:        int32(acc.MinGuests),
		MaxGuests:        int32(acc.MaxGuests),
		Description:      acc.Description,
		AutomaticApprove: acc.AutomaticApprove,
	}
	for _, img := range acc.Pictures {
		accMapped.Pictures = append(accMapped.Pictures, img)
	}
	return accMapped
}

func mapNewAccommodation(newacc *accommodation_service.NewAccommodation) *domain.Accommodation {
	hostId, err := primitive.ObjectIDFromHex(newacc.HostId)
	if err != nil {
		return nil
	}
	bnftsMapped := &domain.Benefits{
		Wifi:        newacc.Benefits.Wifi,
		Kitchen:     newacc.Benefits.Kitchen,
		FreeParking: newacc.Benefits.FreeParking,
	}
	accMapped := &domain.Accommodation{
		HostId:           hostId,
		Name:             newacc.Name,
		Location:         newacc.Location,
		Benefits:         *bnftsMapped,
		MinGuests:        int(newacc.MinGuests),
		MaxGuests:        int(newacc.MaxGuests),
		Description:      newacc.Description,
		AutomaticApprove: newacc.AutomaticApprove,
	}
	for _, img := range newacc.Pictures {
		accMapped.Pictures = append(accMapped.Pictures, img)
	}
	return accMapped
}

func mapUpdateAccommodation(newacc *accommodation_service.Accommodation) *domain.Accommodation {
	hostId, err := primitive.ObjectIDFromHex(newacc.HostId)
	if err != nil {
		return nil
	}
	accId, err := primitive.ObjectIDFromHex(newacc.Id)
	if err != nil {
		return nil
	}
	bnftsMapped := &domain.Benefits{
		Wifi:        newacc.Benefits.Wifi,
		Kitchen:     newacc.Benefits.Kitchen,
		FreeParking: newacc.Benefits.FreeParking,
	}
	accMapped := &domain.Accommodation{
		Id:               accId,
		HostId:           hostId,
		Name:             newacc.Name,
		Location:         newacc.Location,
		Benefits:         *bnftsMapped,
		MinGuests:        int(newacc.MinGuests),
		MaxGuests:        int(newacc.MaxGuests),
		Description:      newacc.Description,
		AutomaticApprove: newacc.AutomaticApprove,
	}
	for _, img := range newacc.Pictures {
		accMapped.Pictures = append(accMapped.Pictures, img)
	}
	return accMapped
}
