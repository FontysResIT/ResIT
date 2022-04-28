package logic

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/FontysResIT/ResIT/model"
	"github.com/stretchr/testify/assert"
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

func (mock *TimeslotRepository) Create(timeslot model.TimeSlot) (model.TimeSlot, error) {
	mock.Called()
	return timeslot, nil
}

func TestGetAllTimeslots(t *testing.T) {
	mockRepo := new(TimeslotRepository)

	timeslot := model.TimeSlot{Id: primitive.NewObjectID(), FromHour: 18, FromMinutes: 30, ToHour: 19, ToMinutes: 30}

	mockRepo.On("All").Return([]model.TimeSlot{timeslot})

	testService := NewTimeSlotLogic(mockRepo)

	testService.GetAllTimeSlots()

	mockRepo.AssertExpectations(t)
}

func TestCreateTimeslot(t *testing.T) {
	mockRepo := new(TimeslotRepository)

	timeslot := model.TimeSlot{FromHour: 18, FromMinutes: 30, ToHour: 19, ToMinutes: 30}
	insertOneResult := new(mongo.InsertOneResult)
	insertOneResult.InsertedID = timeslot.Id

	mockRepo.On("Create").Return(insertOneResult)

	testService := NewTimeSlotLogic(mockRepo)

	result, err := testService.CreateTimeSlot(timeslot)
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, timeslot.Id, result.Id)
	assert.Equal(t, err, nil)
}
