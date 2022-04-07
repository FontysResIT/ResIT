package logic

import (
	"github.com/RealSnowKid/ResIT/model"
	"github.com/RealSnowKid/ResIT/repository"
)

type IDateTimeslotLogic interface {
	GetAllDateTimeslots() []model.DateTimeSlot
}

var _repositoryDT repository.IDateTimeslot

type DTlogic struct{}

func NewDateTimeslotLogic(repository repository.IDateTimeslot) *logic {
	_repositoryDT = repository
	return &logic{}
}

func (*logic) GetAllDateTimeslots() []model.DateTimeSlot {
	return _repositoryDT.All()
}
