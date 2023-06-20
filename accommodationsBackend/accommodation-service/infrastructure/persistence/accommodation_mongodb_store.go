package persistence

import (
	"accommodationsBackend/accommodations-service/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	DATABASE   = "AccomodationsBooking"
	COLLECTION = "accommodations"
)

type AccommodationMongoDBStore struct {
	accommodations *mongo.Collection
}

func NewAccommodationMongoDBStore(client *mongo.Client) domain.AccommodationStore {
	accommodations := client.Database(DATABASE).Collection(COLLECTION)
	return &AccommodationMongoDBStore{
		accommodations: accommodations,
	}
}

func (store *AccommodationMongoDBStore) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AccommodationMongoDBStore) GetAll() ([]*domain.Accommodation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) GetAllProminentAccommodation() ([]*domain.Accommodation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) Insert(accommodation *domain.Accommodation) error {
	result, err := store.accommodations.InsertOne(context.TODO(), accommodation)
	if err != nil {
		return err
	}
	accommodation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}
func (store *AccommodationMongoDBStore) UpdateAccommodation(id string, accommodation *domain.Accommodation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"hostId":           accommodation.HostId,
		"name":             accommodation.Name,
		"location":         accommodation.Location,
		"description":      accommodation.Description,
		"benefits":         accommodation.Benefits,
		"minGuests":        accommodation.MinGuests,
		"maxGuests":        accommodation.MaxGuests,
		"pictures":         accommodation.Pictures,
		"automaticapprove": accommodation.AutomaticApprove,
	}}
	_, err := store.accommodations.UpdateOne(ctx, filter, update)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (store *AccommodationMongoDBStore) GetAccommodationByUserId(id primitive.ObjectID) ([]*domain.Accommodation, error) {
	filter := bson.M{"hostid": id}
	return store.filter(filter)
}
func (store *AccommodationMongoDBStore) GetAccommodationByLocation(location string) ([]*domain.Accommodation, error) {
	filter := bson.M{"location": location}
	return store.filter(filter)
}
func (store *AccommodationMongoDBStore) GetAccommodationByNumberOfGuests(guest int) ([]*domain.Accommodation, error) {
	filter := bson.M{"maxGuests": guest}
	return store.filter(filter)
}
func (store *AccommodationMongoDBStore) GetAccommodationByDateRange(startDate, endDate time.Time) ([]*domain.Accommodation, error) {
	filter := bson.M{
		"start_date": bson.M{
			"$gte": startDate,
		},
		"end_date": bson.M{
			"$lte": endDate,
		},
	}
	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) GetAccommodationByStartDate(date int) ([]*domain.Accommodation, error) {
	filter := bson.M{"": date}
	return store.filter(filter)
}
func (store *AccommodationMongoDBStore) GetAccommodationByEndDate(date int) ([]*domain.Accommodation, error) {
	filter := bson.M{"maxGuests": date}
	return store.filter(filter)
}
func (store *AccommodationMongoDBStore) DeleteAll() {
	store.accommodations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AccommodationMongoDBStore) filter(filter interface{}) ([]*domain.Accommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AccommodationMongoDBStore) filterOne(filter interface{}) (accommodation *domain.Accommodation, err error) {
	result := store.accommodations.FindOne(context.TODO(), filter)
	err = result.Decode(&accommodation)
	return
}

func decode(cursor *mongo.Cursor) (accommodations []*domain.Accommodation, err error) {
	for cursor.Next(context.TODO()) {
		var accommodation domain.Accommodation
		err = cursor.Decode(&accommodation)
		if err != nil {
			return
		}
		accommodations = append(accommodations, &accommodation)
	}
	err = cursor.Err()
	return
}
