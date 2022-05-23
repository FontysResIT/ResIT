package logic

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/FontysResIT/ResIT/model"
	"github.com/stretchr/testify/mock"
)

type DateTimeSlotRepository struct {
	mock.Mock
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

func (mock *DateTimeSlotRepository) ById(id primitive.ObjectID) model.DateTimeSlot {
	args := mock.Called()
	result := args.Get(0)
	return result.(model.DateTimeSlot)
}

func TestGetAllDateTimeSlots(t *testing.T) {
	mockRepo := new(DateTimeSlotRepository)

	dateTimeSlot := model.DateTimeSlot{Id: primitive.NewObjectID(), Date: time.Date(2022, 04, 21, 0, 0, 0, 0, time.FixedZone("CEST", 2*60*60)), Day: time.Thursday, TimeSlot: model.TimeSlot{}}

	mockRepo.On("All").Return([]model.DateTimeSlot{dateTimeSlot})

	testService := NewDateTimeslotLogic(mockRepo)

	testService.GetAllDateTimeslots()

	mockRepo.AssertExpectations(t)
}

func TestAllDTSByDate(t *testing.T) {
	mockRepo := new(DateTimeSlotRepository)

	date := time.Date(2022, 04, 21, 0, 0, 0, 0, time.FixedZone("CEST", 2*60*60))
	dateTimeSlot := model.DateTimeSlot{Id: primitive.NewObjectID(), Date: date, Day: time.Thursday, TimeSlot: model.TimeSlot{}}

	mockRepo.On("AllByDate").Return([]model.DateTimeSlot{dateTimeSlot})

	testService := NewDateTimeslotLogic(mockRepo)

	testService.GetDateTimeslotsByDate(date)

	mockRepo.AssertExpectations(t)
}

func TestAllDTSIdByDate(t *testing.T) {
	mockRepo := new(DateTimeSlotRepository)

	date := time.Date(2022, 04, 21, 0, 0, 0, 0, time.FixedZone("CEST", 2*60*60))

	mockRepo.On("IdByDate").Return([]string{})

	testService := NewDateTimeslotLogic(mockRepo)

	testService.GetDateTimeslotsIdByDate(date)

	mockRepo.AssertExpectations(t)
}
