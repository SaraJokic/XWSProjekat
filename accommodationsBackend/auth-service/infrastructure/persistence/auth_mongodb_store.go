package persistence

import (
	"accommodationsBackend/auth-service/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"

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

func (store *AuthMongoDBStore) ValidateUsernameAndPassword(username string, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//usersCollection := store.getCollection()
	var user domain.User
	err := store.users.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {

		return nil, fmt.Errorf("failed to validate username and password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// password does not match hash
		return nil, fmt.Errorf("failed to validate username and password")
	} else {
		// password matches hash
		return &user, nil
	}

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
func (store *AuthMongoDBStore) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err := store.users.DeleteOne(ctx, filter)
	if err != nil {

		return err
	}

	return nil
}
func (store *AuthMongoDBStore) GetByUsername(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}
