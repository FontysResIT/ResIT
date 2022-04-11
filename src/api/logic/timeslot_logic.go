package logic

import (
	"github.com/RealSnowKid/ResIT/model"
	"github.com/RealSnowKid/ResIT/repository"
)

type ITimeSlotLogic interface {
	GetAllTimeSlots() []model.TimeSlot
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
