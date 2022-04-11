package logic

import (
	"github.com/RealSnowKid/ResIT/model"
	"github.com/RealSnowKid/ResIT/repository"
)

type IDateTimeslotLogic interface {
	GetAllDateTimeslots() []model.DateTimeSlot
	GetDateTimeslotByDate() []model.DateTimeSlot
}

var _repositoryDT repository.IDateTimeslot

type DTlogic struct{}

func NewDateTimeslotLogic(repository repository.IDateTimeslot) *DTlogic {
	_repositoryDT = repository
	return &DTlogic{}
}

func (*DTlogic) GetAllDateTimeslots() []model.DateTimeSlot {
	return _repositoryDT.All()
}

func (*DTlogic) GetDateTimeslotByDate() []model.DateTimeSlot {
	return _repositoryDT.Date()
}
