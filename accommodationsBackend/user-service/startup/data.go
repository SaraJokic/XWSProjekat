package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	var users = []*domain.User{
		{
			Id:      getObjectId("623b0cc3a34d25d8567f9f82"),
			Name:    "User1",
			Surname: "Surname1",
		},
		{
			Id:      getObjectId("623b0cc3a34d25d8567f9f83"),
			Name:    "User2",
			Surname: "Surname2",
		},
	}
*/
func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
