package persistence

import (
	"accommodationsBackend/rating-service/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	DATABASE                     = "AccomodationsAvailability"
	ACCOMMODATIONRATE_COLLECTION = "accommodationratings"
	HOSTRATE_COLLECTION          = "hostratings"
)

type RatingsMongoDBStore struct {
	accommodations *mongo.Collection
	hosts          *mongo.Collection
}

func NewRatingMongoDBStore(client *mongo.Client) domain.RatingsStore {
	accommodations := client.Database(DATABASE).Collection(ACCOMMODATIONRATE_COLLECTION)
	hosts := client.Database(DATABASE).Collection(HOSTRATE_COLLECTION)

	return &RatingsMongoDBStore{
		accommodations: accommodations,
		hosts:          hosts,
	}
}

func (store *RatingsMongoDBStore) GetRateHost(id primitive.ObjectID) (*domain.RateHost, error) {
	filter := bson.M{"_id": id}
	return store.filterOneRateHost(filter)
}

func (store *RatingsMongoDBStore) GetRateAccommodation(id primitive.ObjectID) (*domain.RateAccommodation, error) {
	filter := bson.M{"_id": id}
	return store.filterOneRateAccommodation(filter)
}

func (store *RatingsMongoDBStore) GetAccommodationsRatingsByGuestId(id primitive.ObjectID) ([]*domain.RateAccommodation, error) {
	filter := bson.M{"guestId": id}
	return store.filterAccommodationsRate(filter)
}

func (store *RatingsMongoDBStore) GetHostRatingsByGuestId(id primitive.ObjectID) ([]*domain.RateHost, error) {
	filter := bson.M{"guestId": id}
	return store.filterHostsRate(filter)
}

func (store *RatingsMongoDBStore) GetAccommodationsRatingsByAccommodationId(id primitive.ObjectID) ([]*domain.RateAccommodation, error) {
	filter := bson.M{"accommodationId": id}
	return store.filterAccommodationsRate(filter)
}

func (store *RatingsMongoDBStore) GetHostRatingsByHostId(id primitive.ObjectID) ([]*domain.RateHost, error) {
	filter := bson.M{"hostId": id}
	return store.filterHostsRate(filter)
}

func (store *RatingsMongoDBStore) InsertAccommodationRating(ratingAccommodation *domain.RateAccommodation) error {
	result, err := store.accommodations.InsertOne(context.TODO(), ratingAccommodation)
	if err != nil {
		return err
	}
	ratingAccommodation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *RatingsMongoDBStore) InsertHostRating(ratingHost *domain.RateHost) error {
	result, err := store.hosts.InsertOne(context.TODO(), ratingHost)
	if err != nil {
		return err
	}
	ratingHost.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *RatingsMongoDBStore) UpdateAccommodationRating(id string, rating *domain.RateAccommodation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}

	update := bson.M{"$set": bson.M{
		"guestId":         rating.GuestId,
		"dateRating":      rating.Date,
		"rating":          rating.Rating,
		"accommodationId": rating.AccommodationId,
	}}
	_, err = store.accommodations.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (store *RatingsMongoDBStore) UpdateHostRating(id string, rating *domain.RateHost) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"guestId":    rating.GuestId,
		"dateRating": rating.Date,
		"rating":     rating.Rating,
		"hostId":     rating.HostId,
	}}
	_, err = store.hosts.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (store *RatingsMongoDBStore) DeleteAll() {
	store.accommodations.DeleteMany(context.TODO(), bson.D{{}})
	store.hosts.DeleteMany(context.TODO(), bson.D{{}})
}
func (store *RatingsMongoDBStore) filterOneRateHost(filter interface{}) (hostRating *domain.RateHost, err error) {
	result := store.hosts.FindOne(context.TODO(), filter)
	err = result.Decode(&hostRating)
	return
}
func (store *RatingsMongoDBStore) filterOneRateAccommodation(filter interface{}) (accommodationRating *domain.RateAccommodation, err error) {
	result := store.accommodations.FindOne(context.TODO(), filter)
	err = result.Decode(&accommodationRating)
	return
}
func decodeA(cursor *mongo.Cursor) (accommodations []*domain.RateAccommodation, err error) {
	for cursor.Next(context.TODO()) {
		var a domain.RateAccommodation
		err = cursor.Decode(&a)
		if err != nil {
			return
		}
		accommodations = append(accommodations, &a)
	}
	err = cursor.Err()
	return
}
func decodeH(cursor *mongo.Cursor) (hosts []*domain.RateHost, err error) {
	for cursor.Next(context.TODO()) {
		var a domain.RateHost
		err = cursor.Decode(&a)
		if err != nil {
			return
		}
		hosts = append(hosts, &a)
	}
	err = cursor.Err()
	return
}
func (store *RatingsMongoDBStore) filterAccommodationsRate(filter interface{}) ([]*domain.RateAccommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeA(cursor)
}
func (store *RatingsMongoDBStore) filterHostsRate(filter interface{}) ([]*domain.RateHost, error) {
	cursor, err := store.hosts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeH(cursor)
}
func (store *RatingsMongoDBStore) DeleteRateHost(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	_, err := store.hosts.DeleteOne(ctx, filter)
	if err != nil {

		return err
	}

	return nil
}

func (store *RatingsMongoDBStore) DeleteRateAccommodation(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	_, err := store.accommodations.DeleteOne(ctx, filter)
	if err != nil {

		return err
	}

	return nil
}
