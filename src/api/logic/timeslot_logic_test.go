package logic

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/RealSnowKid/ResIT/model"
	"github.com/stretchr/testify/mock"
)

type TimeslotRepository struct {
	mock.Mock
}

func (mock *TimeslotRepository) All() []model.TimeSlot {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.TimeSlot)
}

func TestGetAllTimeslots(t *testing.T) {
	mockRepo := new(TimeslotRepository)

	timeslot := model.TimeSlot{Id: primitive.NewObjectID(), FromHour: 18, FromMinutes: 30, ToHour: 19, ToMinutes: 30}

	mockRepo.On("All").Return([]model.TimeSlot{timeslot})

	testService := NewTimeSlotLogic(mockRepo)

	testService.GetAllTimeSlots()

	mockRepo.AssertExpectations(t)
}
