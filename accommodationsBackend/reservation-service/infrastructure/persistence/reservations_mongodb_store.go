package persistence

import (
	"accommodationsBackend/reservation-service/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	DATABASE   = "ReservationsStore"
	COLLECTION = "reservations"
)

type ReservationMongoDBStore struct {
	reservations *mongo.Collection
}

func NewReservationMongoDBStore(client *mongo.Client) domain.ReservationStore {
	reservatins := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationMongoDBStore{
		reservations: reservatins,
	}
}

func (store *ReservationMongoDBStore) Get(id primitive.ObjectID) (*domain.Reservation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *ReservationMongoDBStore) GetAll() ([]*domain.Reservation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}
func (store *ReservationMongoDBStore) GetByAccommodationId(id primitive.ObjectID) ([]*domain.Reservation, error) {
	filter := bson.M{"accommodationId": id}
	return store.filter(filter)

}
func (store *ReservationMongoDBStore) GetByUserId(id primitive.ObjectID) ([]*domain.Reservation, error) {
	filter := bson.M{"guestId": id}
	return store.filter(filter)

}

func (store *ReservationMongoDBStore) Insert(reservation *domain.Reservation) error {
	result, err := store.reservations.InsertOne(context.TODO(), reservation)
	if err != nil {
		return err
	}
	reservation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ReservationMongoDBStore) DeleteAll() {
	store.reservations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *ReservationMongoDBStore) filter(filter interface{}) ([]*domain.Reservation, error) {
	cursor, err := store.reservations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *ReservationMongoDBStore) filterOne(filter interface{}) (a *domain.Reservation, err error) {
	result := store.reservations.FindOne(context.TODO(), filter)
	err = result.Decode(&a)
	return
}

func decode(cursor *mongo.Cursor) (reservations []*domain.Reservation, err error) {
	for cursor.Next(context.TODO()) {
		var r domain.Reservation
		err = cursor.Decode(&r)
		if err != nil {
			return
		}
		reservations = append(reservations, &r)
	}
	err = cursor.Err()
	return
}

func (store *ReservationMongoDBStore) Update(id primitive.ObjectID, r *domain.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"guestId":         r.GuestId,
		"accommodationId": r.AccommodationId,
		"startdate":       r.StartDate,
		"enddate":         r.EndDate,
		"numofguests":     r.NumOfGuests,
		"status":          r.Status,
	}}
	_, err := store.reservations.UpdateOne(ctx, filter, update)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (store *ReservationMongoDBStore) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err := store.reservations.DeleteOne(ctx, filter)
	if err != nil {

		return err
	}

	return nil
}
func (store *ReservationMongoDBStore) GetByHostId(id primitive.ObjectID) ([]*domain.Reservation, error) {
	filter := bson.M{"hostId": id}
	return store.filter(filter)
}
