package startup

import (
	"accommodationsBackend/accommodations-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var benefits = domain.Benefits{
	Wifi:        true,
	Kitchen:     true,
	FreeParking: false,
}
var accommodations = []*domain.Accommodation{
	{
		Id:          getObjectId("623b0cc3a34d25d8567f9f855"),
		Name:        "Acc1",
		Location:    "loc1",
		Benefits:    benefits,
		MinGuests:   2,
		MaxGuests:   4,
		Description: "zlatibor",
		HostId:      getObjectId("623a0cc3a34d25d8567f9f855"),
		Pictures:    []string{"/assets/image1.jpg", "/assets/image2.jpg"},
	},
	{
		Id:          getObjectId("623b0cc3a34d25d8547f9f855"),
		Name:        "Acc2",
		Location:    "loc2",
		Benefits:    benefits,
		MinGuests:   1,
		MaxGuests:   2,
		Description: "zlatibor",
		HostId:      getObjectId("623a0cb3a34d25d8567f9f855"),
		Pictures:    []string{"/assets/image3.jpg", "/assets/image4.jpg"},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
