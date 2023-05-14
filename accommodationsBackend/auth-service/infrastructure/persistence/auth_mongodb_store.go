package persistence

import (
	"accommodationsBackend/auth-service/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "AuthBooking"
	COLLECTION = "authUsers"
)

type AuthMongoDBStore struct {
	users *mongo.Collection
}

func NewAuthMongoDBStore(client *mongo.Client) domain.AuthStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &AuthMongoDBStore{
		users: users,
	}
}
func (store *AuthMongoDBStore) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AuthMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AuthMongoDBStore) Insert(user *domain.User) error {
	user.Id = primitive.NewObjectID()

	result, err := store.users.InsertOne(context.TODO(), user)

	fmt.Println("usao sam u inser funjciju")
	fmt.Println("ovo je user kog upisujem u bazu ", user)
	if err != nil {
		return err
	}

	user.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AuthMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AuthMongoDBStore) filterOne(filter interface{}) (user *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func decode(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var user domain.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}
