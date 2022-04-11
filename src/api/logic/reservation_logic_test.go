package logic

import (
	"fmt"
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

func (mock *MockRepository) All() []model.Reservation {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Reservation)
}
func (mock *MockRepository) Create(reservation model.Reservation) *mongo.InsertOneResult {
	args := mock.Called()
	result := args.Get(0)
	return result.(*mongo.InsertOneResult)
}
func TestGetAll(t *testing.T) {
	mockRepo := new(MockRepository)

	userProfile := model.Reservation{Id: primitive.ObjectID{}, FirstName: "Peter", LastName: "Pancakes", DateTimeSlotId: "0", Email: "peter@example.com", GuestCount: 2, PhoneNumber: "+31 6 12345678"}
	// Setup expectations
	mockRepo.On("All").Return([]model.Reservation{userProfile})

	testService := NewReservationLogic(mockRepo)

	result := testService.GetAllReservations()
	fmt.Println(result)
	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, userProfile.Id, result[0].Id)
	assert.Equal(t, userProfile.Email, result[0].Email)
	assert.Equal(t, userProfile.GuestCount, result[0].GuestCount)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	userProfile := model.Reservation{Id: primitive.ObjectID{}, FirstName: "Peter", LastName: "Pancakes", DateTimeSlotId: "0", Email: "peter@example.com", GuestCount: 2, PhoneNumber: "+31 6 12345678"}
	insertOneResult := new(mongo.InsertOneResult)
	insertOneResult.InsertedID = userProfile.Id
	// Setup expectations
	mockRepo.On("Create").Return(insertOneResult)

	testService := NewReservationLogic(mockRepo)

	result := testService.CreateReservation(userProfile)
	fmt.Println(result)	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, userProfile.Id, result.InsertedID)
}
