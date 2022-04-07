package logic

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"

	"github.com/RealSnowKid/ResIT/model"
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

func TestGetAll(t *testing.T) {
	mockRepo := new(MockRepository)

	userProfile := model.Reservation{Id: primitive.NewObjectID(), FirstName: "Peter", LastName: "Pancakes", DateTimeSlotId: primitive.NewObjectID(), Email: "peter@example.com", GuestCount: 2, PhoneNumber: "+31 6 12345678"}

	mockRepo.On("All").Return([]model.Reservation{userProfile})

	testService := NewReservationLogic(mockRepo)

	testService.GetAllReservations()

	mockRepo.AssertExpectations(t)
}
