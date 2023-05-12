package persistence

import (
	"accommodationsBackend/user-service/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	DATABASE   = "AccomodationsBooking"
	COLLECTION = "users"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) Get(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) Register(user *domain.User) error {
	user.Id = primitive.NewObjectID()

	result, err := store.users.InsertOne(context.TODO(), user)

	fmt.Println("usao sam u registermongoDB funjciju")
	fmt.Println("ovo je user kog upisujem u bazu ", user)
	if err != nil {
		return err
	}

	user.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserMongoDBStore) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *UserMongoDBStore) CheckIfEmailAndUsernameExist(email string, username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("ovo je username i email ", username, email)
	fmt.Println("kontekst ", ctx)
	//usersCollection := store.getCollection()
	var userEmail domain.User
	err := store.users.FindOne(ctx, bson.M{"email": email}).Decode(&userEmail)
	fmt.Println("ovo je userEmail", userEmail)
	if err == nil {
		// email exists in database
		return true, nil
	} else if err != mongo.ErrNoDocuments {
		return false, err
	}

	var userUsername domain.User
	err = store.users.FindOne(ctx, bson.M{"username": username}).Decode(&userUsername)
	fmt.Println("ovo je userUsername", userUsername)
	if err == nil {
		// username exists in database

		return true, nil
	} else if err != mongo.ErrNoDocuments {
		return false, err
	}

	// email and username aren't in the database
	return false, nil
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (user *domain.User, err error) {
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
