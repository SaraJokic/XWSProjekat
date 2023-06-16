package persistence

import (
	"accommodationsBackend/availability-service/domain"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	DATABASE   = "AccomodationsAvailability"
	COLLECTION = "availabilities"
)

type AvailabilityMongoDBStore struct {
	availabilities *mongo.Collection
}

func NewAvailabilityMongoDBStore(client *mongo.Client) domain.AvailabilityStore {
	availabilities := client.Database(DATABASE).Collection(COLLECTION)
	return &AvailabilityMongoDBStore{
		availabilities: availabilities,
	}
}

func (store *AvailabilityMongoDBStore) Get(id primitive.ObjectID) (*domain.Availability, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AvailabilityMongoDBStore) GetAll() ([]*domain.Availability, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}
func (store *AvailabilityMongoDBStore) GetByAccommodationId(id primitive.ObjectID) (*domain.Availability, error) {
	filter := bson.M{"accommodationId": id}
	return store.filterOne(filter)

}

func (store *AvailabilityMongoDBStore) Insert(availability *domain.Availability) error {
	result, err := store.availabilities.InsertOne(context.TODO(), availability)
	if err != nil {
		return err
	}
	availability.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AvailabilityMongoDBStore) DeleteAll() {
	store.availabilities.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AvailabilityMongoDBStore) filter(filter interface{}) ([]*domain.Availability, error) {
	cursor, err := store.availabilities.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AvailabilityMongoDBStore) filterOne(filter interface{}) (a *domain.Availability, err error) {
	result := store.availabilities.FindOne(context.TODO(), filter)
	err = result.Decode(&a)
	return
}
func (store *AvailabilityMongoDBStore) FindAvailabilitySlotsByDateRange(startDate time.Time, endDate time.Time) ([]domain.AvailabilitySlot, error) {
	allAvailability, err := store.GetAll()
	if err != nil {
		return nil, err
	}

	var slots []domain.AvailabilitySlot

	for _, availability := range allAvailability {
		for _, slot := range availability.AvailableSlots {
			if slot.StartDate.After(startDate) && slot.EndDate.Before(endDate) {
				slots = append(slots, slot)
			}
		}
	}

	return slots, nil
}

func decode(cursor *mongo.Cursor) (availabilities []*domain.Availability, err error) {
	for cursor.Next(context.TODO()) {
		var a domain.Availability
		err = cursor.Decode(&a)
		if err != nil {
			return
		}
		availabilities = append(availabilities, &a)
	}
	err = cursor.Err()
	return
}

func (store *AvailabilityMongoDBStore) Update(id primitive.ObjectID, a *domain.Availability) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//fmt.Println("ovo je user koji je stigao u repo", user)
	fmt.Println("ovo je novi availability ", a.ChangePrice)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"accommodationId":  a.AccommodationId,
		"availability":     a.AvailableSlots,
		"price":            a.Price,
		"ispriceperperson": a.IsPricePerGuest,
		"changePrice":      a.ChangePrice,
	}}
	_, err := store.availabilities.UpdateOne(ctx, filter, update)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
