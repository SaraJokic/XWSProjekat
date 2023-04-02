package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"xwsproj/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment
func NewUserRepo(ctx context.Context, logger *log.Logger) (*UserRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &UserRepo{
		cli:    client,
		logger: logger,
	}, nil
}
func (pr *UserRepo) GetAll() (model.Users, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := pr.getCollection()

	var users model.Users
	usersCursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &users); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return users, nil
}

func (pr *UserRepo) GetById(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := pr.getCollection()

	var user model.User
	objID, _ := primitive.ObjectIDFromHex(id)
	err := usersCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return &user, nil
}

func (pr *UserRepo) FindByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := pr.getCollection()

	var user model.User
	err := usersCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepo) CheckIfEmailAndUsernameExist(email string, username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.getCollection()
	var userEmail model.User
	err := usersCollection.FindOne(ctx, bson.M{"email": email}).Decode(&userEmail)
	if err == nil {
		// email exists in database
		return true, nil
	} else if err != mongo.ErrNoDocuments {
		return false, err
	}

	var userUsername model.User
	err = usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&userUsername)
	if err == nil {
		// username exists in database
		return true, nil
	} else if err != mongo.ErrNoDocuments {
		return false, err
	}

	// email and username aren't in the database
	return false, nil
}

func (ur *UserRepo) ValidateUsernameAndPassword(username string, password string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.getCollection()
	var user model.User
	err := usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// password does not match hash
		return nil, err
	} else {
		// password matches hash
		return &user, nil
	}

}

func (pr *UserRepo) FindByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := pr.getCollection()

	var user model.User
	err := usersCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}

	return &user, nil
}

func (pr *UserRepo) Insert(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := pr.getCollection()

	result, err := usersCollection.InsertOne(ctx, &user)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (pr *UserRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := usersCollection.DeleteOne(ctx, filter)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

// Check database connection
func (fr *UserRepo) PingUserRepo() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := fr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		fr.logger.Println(err)
	}

	// Print available databases
	databases, err := fr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fr.logger.Println(err)
	}
	fmt.Println(databases)
}

// Disconnect from database
func (fr *UserRepo) DisconnectUserRepo(ctx context.Context) error {
	err := fr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (pr *UserRepo) getCollection() *mongo.Collection {
	userDatabase := pr.cli.Database("Airport")
	userCollection := userDatabase.Collection("RegisteredUsers")
	return userCollection
}
