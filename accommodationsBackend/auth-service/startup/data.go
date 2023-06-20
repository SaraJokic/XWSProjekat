package startup

import (
	"accommodationsBackend/auth-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.User{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f82"),
		Username: "User1",
		Password: "123",
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
		Username: "User2",
		Password: "123",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
