package logic

import (
	"time"

	"github.com/FontysResIT/ResIT/model"
	"github.com/FontysResIT/ResIT/repository"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDateTimeslotLogic interface {
	GetAllDateTimeslots() []model.DateTimeSlot
	GetDateTimeslotsByDate(time.Time) []model.DateTimeSlot
	GetDateTimeslotsIdByDate(time.Time) []string
	GetDateTimeslotById(id string) model.DateTimeSlot
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

func (*DTlogic) GetDateTimeslotsIdByDate(param time.Time) []string {
	return _repositoryDT.IdByDate(param)
}

func (*DTlogic) GetDateTimeslotById(id string) model.DateTimeSlot {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
	}
	return _repositoryDT.ById(_id)
}
