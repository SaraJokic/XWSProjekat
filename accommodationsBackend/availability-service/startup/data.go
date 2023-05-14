package startup

import (
	"accommodationsBackend/availability-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var availableSlots = []domain.AvailabilitySlot{
	{
		SlotId:    getObjectId("663b0cc3a34d25d8567t9f82"),
		StartDate: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 7, 31, 0, 0, 0, 0, time.UTC),
	},
	{
		SlotId:    getObjectId("623b0cc3a34d25d8567t9f83"),
		StartDate: time.Date(2023, 5, 4, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 5, 10, 0, 0, 0, 0, time.UTC),
	},
}
var priceChanges = []domain.PriceChange{
	{
		StartDate: time.Date(2023, 7, 5, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 7, 20, 0, 0, 0, 0, time.UTC),
		Change:    40,
	},
	{
		StartDate: time.Date(2023, 5, 4, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 5, 10, 0, 0, 0, 0, time.UTC),
		Change:    10,
	},
}

var availabilities = []*domain.Availability{
	{
		Id:              getObjectId("623b0cc3a34d25d8567t9f82"),
		AccommodationId: getObjectId("623b0cc3a34d25d8567f9f855"),
		AvailableSlots:  availableSlots,
		Price:           120.0,
		IsPricePerGuest: true,
		ChangePrice:     priceChanges,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
