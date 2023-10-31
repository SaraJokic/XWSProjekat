package application

import (
	domain "accommodationsBackend/availability-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type AvailabilityService struct {
	store domain.AvailabilityStore
}

func NewAvailabilityService(store domain.AvailabilityStore) *AvailabilityService {
	return &AvailabilityService{
		store: store,
	}
}

func (service *AvailabilityService) Get(id primitive.ObjectID) (*domain.Availability, error) {
	return service.store.Get(id)
}

func (service *AvailabilityService) GetAll() ([]*domain.Availability, error) {
	return service.store.GetAll()
}
func (service *AvailabilityService) Update(id primitive.ObjectID, availability *domain.Availability) error {
	return service.store.Update(id, availability)
}
func (service *AvailabilityService) GetByAccommodationId(id primitive.ObjectID) (*domain.Availability, error) {
	return service.store.GetByAccommodationId(id)

}
func (service *AvailabilityService) FindAvailabilitySlotsByDateRange(startDate time.Time, endDate time.Time) ([]domain.AvailabilitySlot, error) {
	return service.store.FindAvailabilitySlotsByDateRange(startDate, endDate)
}

func (service *AvailabilityService) Insert(availability *domain.Availability) error {
	if availability.Id.IsZero() {
		availability.Id = primitive.NewObjectID()
	}
	return service.store.Insert(availability)
}
func (service *AvailabilityService) updateAvailableSlots(id primitive.ObjectID, selectedStartDate time.Time, selectedEndDate time.Time) error {
	availability, _ := service.GetByAccommodationId(id)
	updatedSlots := make([]domain.AvailabilitySlot, 0)
	slots := availability.AvailableSlots
	for _, slot := range slots {
		// Check if the slot overlaps with the selected dates
		if slot.StartDate.Before(selectedStartDate) && slot.EndDate.After(selectedEndDate) {
			// Split the slot into two parts: before the selected start date and after the selected end date
			// Add the first part to the updated slots
			updatedSlots = append(updatedSlots, domain.AvailabilitySlot{StartDate: slot.StartDate, EndDate: selectedStartDate})

			// Add the second part to the updated slots
			updatedSlots = append(updatedSlots, domain.AvailabilitySlot{StartDate: selectedEndDate, EndDate: slot.EndDate})
		} else if slot.StartDate.Before(selectedStartDate) && slot.EndDate.After(selectedStartDate) {
			// Adjust the end date of the slot to be the selected start date
			slot.EndDate = selectedStartDate
			updatedSlots = append(updatedSlots, slot)
		} else if slot.StartDate.Before(selectedEndDate) && slot.EndDate.After(selectedEndDate) {
			// Adjust the start date of the slot to be the selected end date
			slot.StartDate = selectedEndDate
			updatedSlots = append(updatedSlots, slot)
		} else if slot.StartDate.Equal(selectedStartDate) && slot.EndDate.After(selectedEndDate) {

			updatedSlots = append(updatedSlots, domain.AvailabilitySlot{StartDate: selectedEndDate, EndDate: slot.EndDate})

		} else if slot.StartDate.Before(selectedStartDate) && slot.EndDate.Equal(selectedEndDate) {

			updatedSlots = append(updatedSlots, domain.AvailabilitySlot{StartDate: slot.StartDate, EndDate: selectedStartDate})

		} else if slot.StartDate.Equal(selectedStartDate) && slot.EndDate.Equal(selectedEndDate) {
			continue
		} else {
			// No overlap, keep the slot as it is
			updatedSlots = append(updatedSlots, slot)
		}
	}
	availability.AvailableSlots = updatedSlots
	// Zameniti da su updatedSlots ustv availabilit.AvailabilitySlots
	err := service.Update(availability.Id, availability)
	if err != nil {
		return err
	}
	return nil
}
func (service *AvailabilityService) MakeSlotAvailable(id string, startDate string, endDate string) (domain.Availability, error) {
	layout := "2006-01-02T15:04:05Z"
	startdate, err := time.Parse(layout, startDate)
	if err != nil {
		log.Println("Failed to parse the string of StartDate: ", err)
		return domain.Availability{}, err
	}
	enddate, err := time.Parse(layout, endDate)
	if err != nil {
		log.Println("Failed to parse the string of EndDate: ", err)
		return domain.Availability{}, err
	}
	accId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Availability{}, err
	}
	newAvailability, _ := service.store.MakeSlotAvailable(accId, startdate, enddate)
	return *newAvailability, nil
}
