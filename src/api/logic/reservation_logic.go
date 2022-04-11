package logic

import (
	"time"

	"github.com/RealSnowKid/ResIT/model"
	"github.com/RealSnowKid/ResIT/repository"
)

type IReservationLogic interface {
	GetAllReservations() []model.Reservation
	GetAllReservationsByDate(time.Time) []model.Reservation
}

var _repository repository.IReservation
var _DTSrepository repository.IDateTimeslot

type logic struct{}

func NewReservationLogic(repository repository.IReservation, dtsRepository repository.IDateTimeslot) *logic {
	_repository = repository
	_DTSrepository = dtsRepository
	return &logic{}
}

func (*logic) GetAllReservations() []model.Reservation {
	return _repository.All()
}

func (*logic) GetAllReservationsByDate(date time.Time) []model.Reservation {
	return _repository.AllByDate(_DTSrepository.IdByDate(date))
}
