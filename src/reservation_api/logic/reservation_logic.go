package logic

import (
	"time"

	"github.com/FontysResIT/ResIT/model"
	"github.com/FontysResIT/ResIT/repository"
	"github.com/FontysResIT/ResIT/util"
)

type IReservationLogic interface {
	GetAllReservations() []model.Reservation
	GetAllReservationsByDate(time.Time) []model.Reservation
	CreateReservation(model.Reservation) (model.Reservation, error)
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

func (*logic) CreateReservation(reservation model.Reservation) (model.Reservation, error) {
	result, err := _repository.Create(reservation)
	if err == nil {
		go util.CreateReservation(result)
	}
	return result, err

}
