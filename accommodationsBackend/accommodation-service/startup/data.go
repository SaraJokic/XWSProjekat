package startup

import (
	"accommodationsBackend/accommodation-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var accommodations = []*domain.Accommodation{
	{
		Id:      getObjectId("623b0cc3a34d25d8567f9f855"),
		Name:    "Acc1",
		Location: "loc1",
		Benefits: "wifi"
	},

	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Location  string             `bson:"surname"`
	Benefits  Benefits           `bson:"benefits"`
	MinGuests int                `bson:"minGuests"`
	MaxGuests int                `bson:"maxGuests"`


	{
		Id:      getObjectId("623b0cc3a34d25d8567f9f83"),
		Name:    "User2",
		Surname: "Surname2",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
