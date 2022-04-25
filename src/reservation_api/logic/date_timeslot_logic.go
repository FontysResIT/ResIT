package logic

import (
	"time"

	"github.com/FontysResIT/ResIT/model"
	"github.com/FontysResIT/ResIT/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDateTimeslotLogic interface {
	GetAllDateTimeslots() []model.DateTimeSlot
	GetDateTimeslotsByDate(time.Time) []model.DateTimeSlot
	GetDateTimeslotsIdByDate(time.Time) []primitive.ObjectID
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

func (*DTlogic) GetDateTimeslotsByDate(param time.Time) []model.DateTimeSlot {
	return _repositoryDT.AllByDate(param)
}

func (*DTlogic) GetDateTimeslotsIdByDate(param time.Time) []primitive.ObjectID {
	return _repositoryDT.IdByDate(param)
}
