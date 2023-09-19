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
		Name:        "Apartman King",
		Location:    "Zlatibor",
		Benefits:    benefits,
		MinGuests:   2,
		MaxGuests:   4,
		Description: "Charming and cozy apartment nestled in the serene landscapes of Zlatibor. This modern haven comes fully equipped with all amenities, offering a perfect blend of comfort and tradition. Step out to enjoy the pristine nature, or relax on the balcony with breathtaking mountain views. Located close to popular attractions and local eateries, it's an ideal getaway for those looking to experience the true essence of Zlatibor",
		HostId:      getObjectId("623a0cc3a34d25d8567f9f855"),
		Pictures:    []string{"/assets/king1.jpg", "/assets/king2.jpg", "/assets/king3.jpg"},
	},
	{
		Id:          getObjectId("623b0cc3a34d25d8547f9f855"),
		Name:        "Hotel Buket",
		Location:    "Zlatibor",
		Benefits:    benefits,
		MinGuests:   1,
		MaxGuests:   2,
		Description: "Elegant hotel in Zlatibor offering modern comforts amidst stunning mountain vistas. Conveniently located near attractions, it's the perfect mountain retreat",
		HostId:      getObjectId("623a0cb3a34d25d8567f9f855"),
		Pictures:    []string{"/assets/buket1.jpg", "/assets/buket2.jpg", "/assets/buket3.jpg", "/assets/buket4.jpg"},
	},
	{
		Id:          getObjectId("623b0cc3a38d25d8547f9f855"),
		Name:        "Hotel Sun Island Maldives",
		Location:    "Male",
		Benefits:    benefits,
		MinGuests:   1,
		MaxGuests:   2,
		Description: "Experience paradise at our luxury hotel in the Maldives. Nestled amidst azure waters and pristine white sands, it's a tropical haven offering unparalleled ocean views, world-class amenities, and a serene escape from the everyday.",
		HostId:      getObjectId("623a0cb3a34d25d8567f9f855"),
		Pictures:    []string{"/assets/sun1.jpg", "/assets/sun2.jpg", "/assets/sun3.jpg"},
	},
	{
		Id:          getObjectId("623b0cc3a38d25d8547f9f855"),
		Name:        "Asterte suites",
		Location:    "Santorini",
		Benefits:    benefits,
		MinGuests:   1,
		MaxGuests:   2,
		Description: "Immerse yourself in Grecian elegance at our Santorini retreat. Perched atop stunning cliffs, our hotel boasts breathtaking caldera views, sun-kissed terraces, and sophisticated Cycladic charm.",
		HostId:      getObjectId("623a0cb3a34d25d8567f9f855"),
		Pictures:    []string{"/assets/astarte1.jpg", "/assets/astarte2.jpg", "/assets/astarte3.jpg"},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
