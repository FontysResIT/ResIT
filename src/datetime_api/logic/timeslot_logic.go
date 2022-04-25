package logic

import (
	"github.com/FontysResIT/datetime_api/model"
	"github.com/FontysResIT/datetime_api/repository"
)

type ITimeSlotLogic interface {
	GetAllTimeSlots() []model.TimeSlot
	CreateTimeSlot(model.TimeSlot) (model.TimeSlot, error)
}

var _repositoryTS repository.ITimeSlot

type TSlogic struct{}

func NewTimeSlotLogic(repository repository.ITimeSlot) *TSlogic {
	_repositoryTS = repository
	return &TSlogic{}
}

func (*TSlogic) GetAllTimeSlots() []model.TimeSlot {
	return _repositoryTS.All()
}

func (*TSlogic) CreateTimeSlot(timeslot model.TimeSlot) (model.TimeSlot, error) {
	ts, err := _repositoryTS.Create(timeslot)
	return ts, err
}
