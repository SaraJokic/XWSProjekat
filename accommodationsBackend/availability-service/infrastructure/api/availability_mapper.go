package api

import (
	"accommodationsBackend/availability-service/domain"
	availability_service "accommodationsBackend/common/proto/availability-service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func mapAvailability(a *domain.Availability) *availability_service.Availability {
	availabilityMapped := &availability_service.Availability{
		Id:               a.Id.Hex(),
		AccommodationId:  a.AccommodationId.Hex(),
		Price:            float32(a.Price),
		IsPricePerPerson: a.IsPricePerGuest,
	}
	for _, slot := range a.AvailableSlots {
		availabilityMapped.AvailableSlots = append(availabilityMapped.AvailableSlots,
			&availability_service.AvailableSlots{
				SlotId:    slot.SlotId.Hex(),
				StartDate: slot.StartDate.String(),
				EndDate:   slot.EndDate.String(),
			})
	}
	for _, prange := range a.ChangePrice {
		availabilityMapped.ChangePrice = append(availabilityMapped.ChangePrice,
			&availability_service.PriceChange{
				Startdate: prange.StartDate.String(),
				Enddate:   prange.EndDate.String(),
				Change:    float32(prange.Change),
			})
	}
	return availabilityMapped
}
func mapDomainAvailability(a *availability_service.CreateNewAvailabilityRequest) *domain.Availability {
	accId, err := primitive.ObjectIDFromHex(a.AccommodationId)
	if err != nil {
		return nil
	}
	availabilityMapped := &domain.Availability{
		AccommodationId: accId,
		Price:           float64(a.Price),
		IsPricePerGuest: a.IsPricePerPerson,
	}

	for _, slot := range a.AvailableSlots {
		layout := "2006-01-02T15:04:05Z"
		//fmt.Println("Start datum price change", request.PriceChange.Startdate)
		startdate, err := time.Parse(layout, slot.StartDate)
		if err != nil {
			log.Println("Failed to parse the string of StartDate: ", err)
			return nil
		}
		enddate, err := time.Parse(layout, slot.EndDate)
		if err != nil {
			log.Println("Failed to parse the string of EndDate: ", err)
			return nil
		}
		availabilityMapped.AvailableSlots = append(availabilityMapped.AvailableSlots,
			domain.AvailabilitySlot{
				SlotId:    primitive.NewObjectID(),
				StartDate: startdate,
				EndDate:   enddate,
			})
	}
	for _, prange := range a.ChangePrice {
		layout := "2006-01-02T15:04:05Z"
		//fmt.Println("Start datum price change", request.PriceChange.Startdate)
		startdate, err := time.Parse(layout, prange.Startdate)
		if err != nil {
			log.Println("Failed to parse the string of StartDate: ", err)
			return nil
		}
		enddate, err := time.Parse(layout, prange.Enddate)
		if err != nil {
			log.Println("Failed to parse the string of EndDate: ", err)
			return nil
		}
		availabilityMapped.ChangePrice = append(availabilityMapped.ChangePrice,
			domain.PriceChange{
				StartDate: startdate,
				EndDate:   enddate,
				Change:    float64(prange.Change),
			})
	}
	return availabilityMapped
}
