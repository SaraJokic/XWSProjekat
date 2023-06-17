package persistence

import (
	"accommodationsBackend/rating-service/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	//TODO implement me
	panic("implement me")
}

func (store *RatingsMongoDBStore) GetHostRatingsByGuestId(id primitive.ObjectID) ([]*domain.RateHost, error) {
	//TODO implement me
	panic("implement me")
}

func (store *RatingsMongoDBStore) GetAccommodationsRatingsByAccommodationId(id primitive.ObjectID) ([]*domain.RateAccommodation, error) {
	//TODO implement me
	panic("implement me")
}

func (store *RatingsMongoDBStore) GetHostRatingsByHostId(id primitive.ObjectID) ([]*domain.RateHost, error) {
	//TODO implement me
	panic("implement me")
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
	result, err := store.accommodations.InsertOne(context.TODO(), ratingHost)
	if err != nil {
		return err
	}
	ratingHost.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *RatingsMongoDBStore) UpdateAccommodationRating(id string, rating *domain.RateAccommodation) error {
	//TODO implement me
	panic("implement me")
}

func (store *RatingsMongoDBStore) UpdateHostRating(id string, rating *domain.RateHost) error {
	//TODO implement me
	panic("implement me")
}

func (store *RatingsMongoDBStore) DeleteAll() {
	//TODO implement me
	panic("implement me")
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
