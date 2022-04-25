package logic

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/RealSnowKid/ResIT/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ReservationRepository struct {
	mock.Mock
}
type MockProducer struct {
	mock.Mock
}

var userProfile = model.Reservation{Id: primitive.NewObjectID(), FirstName: "Peter", LastName: "Pancakes", DateTimeSlotId: primitive.NewObjectID(), Email: "peter@example.com", GuestCount: 2, PhoneNumber: "+31 6 12345678"}

func (mock *ReservationRepository) All() []model.Reservation {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Reservation)
}
func (mock *ReservationRepository) Create(reservation model.Reservation) (model.Reservation, error) {
	mock.Called()
	return reservation, nil
}
func (mock *ReservationRepository) AllByDate([]string) []model.Reservation {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Reservation)
}

func (mock *DateTimeSlotRepository) All() []model.DateTimeSlot {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.DateTimeSlot)
}
func (mock *DateTimeSlotRepository) AllByDate(param time.Time) []model.DateTimeSlot {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.DateTimeSlot)
}
func (mock *DateTimeSlotRepository) IdByDate(param time.Time) []string {
	args := mock.Called()
	result := args.Get(0)
	return result.([]string)
}
func (mock *ReservationRepository) Cancel(id string) (model.Reservation, error) {
	mock.Called()
	reservation := userProfile
	reservation.IsCancelled = true
	return reservation, nil
}

func TestGetAll(t *testing.T) {
	mockRepo := new(ReservationRepository)
	mockRepo2 := new(DateTimeSlotRepository)

	userProfile := model.Reservation{Id: primitive.NewObjectID(), FirstName: "Peter", LastName: "Pancakes", DateTimeSlotId: primitive.NewObjectID(), Email: "peter@example.com", GuestCount: 2, PhoneNumber: "+31 6 12345678"}

	mockRepo.On("All").Return([]model.Reservation{userProfile})

	testService := NewReservationLogic(mockRepo, mockRepo2)

	testService.GetAllReservations()

	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(ReservationRepository)
	mockRepo2 := new(DateTimeSlotRepository)

	userProfile := model.Reservation{FirstName: "Peter", LastName: "Pancakes", DateTimeSlotId: primitive.NewObjectID(), Email: "peter@example.com", GuestCount: 2, PhoneNumber: "+31 6 12345678"}

	// Setup expectations
	mockRepo.On("Create").Return(userProfile)

	testService := NewReservationLogic(mockRepo, mockRepo2)

	result, err := testService.CreateReservation(userProfile)
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, userProfile.Id, result.Id)
	assert.Equal(t, err, nil)
}

func TestCancel(t *testing.T) {
	mockRepo := new(ReservationRepository)
	mockRepo2 := new(DateTimeSlotRepository)
	// Setup expectations
	mockRepo.On("Cancel").Return(userProfile)

	testService := NewReservationLogic(mockRepo, mockRepo2)

	result, err := testService.CancelReservation(userProfile.Id.Hex())
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, userProfile.Id, result.Id)
	assert.NotEqual(t, userProfile.IsCancelled, result.IsCancelled)
	assert.Equal(t, result.IsCancelled, true)
	assert.Equal(t, err, nil)
}
func TestGetAllRByDate(t *testing.T) {
	mockRepo := new(ReservationRepository)
	mockRepo2 := new(DateTimeSlotRepository)

	testDTSId := primitive.NewObjectID()
	testRes := model.Reservation{Id: primitive.NewObjectID(), FirstName: "Peter", LastName: "Pancakes", DateTimeSlotId: testDTSId, Email: "peter@example.com", GuestCount: 2, PhoneNumber: "+31 6 12345678"}
	testDTS := model.DateTimeSlot{Id: testDTSId, Date: time.Date(2022, 04, 21, 0, 0, 0, 0, time.FixedZone("CEST", 2*60*60)), Day: time.Thursday, TimeSlot: model.TimeSlot{}}

	mockRepo.On("AllByDate").Return([]model.Reservation{testRes})
	mockRepo2.On("IdByDate").Return([]primitive.ObjectID{testDTSId})

	testService := NewReservationLogic(mockRepo, mockRepo2)

	testService.GetAllReservationsByDate(testDTS.Date)

	mockRepo.AssertExpectations(t)
}
