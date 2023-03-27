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
)

// NoSQL: ProductRepo struct encapsulating Mongo api client
type FlightRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment
func NewFlightRepo(ctx context.Context, logger *log.Logger) (*FlightRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &FlightRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (fr *FlightRepo) DisconnectFlightRepo(ctx context.Context) error {
	err := fr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (fr *FlightRepo) PingFlightRepo() {
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

func (pr *FlightRepo) GetAllFlights() (model.Flights, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := pr.getCollection()

	var flights model.Flights
	flightsCursor, err := flightsCollection.Find(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return flights, nil
}

func (pr *FlightRepo) GetFlightById(id string) (*model.Flight, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := pr.getCollection()

	var flight model.Flight
	objID, _ := primitive.ObjectIDFromHex(id)
	err := flightsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&flight)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return &flight, nil
}

func (pr *FlightRepo) CreateFlight(flight *model.Flight) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := pr.getCollection()

	result, err := flightsCollection.InsertOne(ctx, &flight)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (pr *FlightRepo) DeleteFlight(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := flightsCollection.DeleteOne(ctx, filter)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}
func (pr *FlightRepo) getCollection() *mongo.Collection {
	flightDatabase := pr.cli.Database("Airport")
	flightsCollection := flightDatabase.Collection("flights")
	return flightsCollection
}
