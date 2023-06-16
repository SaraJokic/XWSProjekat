package startup

import (
	"accommodationsBackend/user-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var users = []*domain.User{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f82"),
		Name:     "User1",
		LastName: "Surname1",
		Username: "user1",
		Email:    "user1@gmail.com",
		Password: "1234",
		City:     "Sabac",
		Country:  "Serbia",
		Role:     1,
	},
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f83"),
		Name:     "User2",
		LastName: "Surname2",
		Username: "user2",
		Email:    "user2@gmail.com",
		Password: "1234",
		City:     "Novi Sad",
		Country:  "Serbia",
		Role:     1,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
