package logic

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/RealSnowKid/ResIT/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

type MockProducer struct {
	mock.Mock
}

func (mock *MockProducer) CreateReservation(reservation model.Reservation) {
	mock.Called()

}

func (mock *MockRepository) All() []model.Reservation {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Reservation)
}
func (mock *MockRepository) Create(reservation model.Reservation) (model.Reservation, error) {
	mock.Called()
	return reservation, nil
}
func TestGetAll(t *testing.T) {
	mockRepo := new(MockRepository)

	userProfile := model.Reservation{Id: primitive.NewObjectID(), FirstName: "Peter", LastName: "Pancakes", DateTimeSlotId: primitive.NewObjectID(), Email: "peter@example.com", GuestCount: 2, PhoneNumber: "+31 6 12345678"}

	mockRepo.On("All").Return([]model.Reservation{userProfile})

	testService := NewReservationLogic(mockRepo)

	testService.GetAllReservations()

	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	userProfile := model.Reservation{FirstName: "Peter", LastName: "Pancakes", DateTimeSlotId: primitive.NewObjectID(), Email: "peter@example.com", GuestCount: 2, PhoneNumber: "+31 6 12345678"}
	insertOneResult := new(mongo.InsertOneResult)
	insertOneResult.InsertedID = userProfile.Id
	// Setup expectations
	mockRepo.On("Create").Return(insertOneResult)

	testService := NewReservationLogic(mockRepo)

	result, err := testService.CreateReservation(userProfile)
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, userProfile.Id, result.Id)
	assert.Equal(t, err, nil)
}
