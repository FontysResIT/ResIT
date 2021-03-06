package logic

import (
	"time"

	"github.com/FontysResIT/ResIT/model"
	"github.com/FontysResIT/ResIT/repository"
	"github.com/FontysResIT/ResIT/util"
)

type IReservationLogic interface {
	GetAllReservations() []model.ReservationReadDTO
	GetAllReservationsByDate(time.Time) []model.Reservation
	CreateReservation(model.Reservation) (model.Reservation, error)
	CancelReservation(string) (model.Reservation, error)
}

var _repository repository.IReservation
var _DTSrepository repository.IDateTimeslot

type logic struct{}

func NewReservationLogic(repository repository.IReservation, dtsRepository repository.IDateTimeslot) *logic {
	_repository = repository
	_DTSrepository = dtsRepository
	return &logic{}
}

func (*logic) GetAllReservations() []model.ReservationReadDTO {
	return _repository.All()
}

func (*logic) GetAllReservationsByDate(date time.Time) []model.Reservation {
	return _repository.AllByDate(_DTSrepository.IdByDate(date))
}

func (*logic) CreateReservation(reservation model.Reservation) (model.Reservation, error) {
	result, err := _repository.Create(reservation)
	if err == nil {
		go func() {
			r := _repository.GetById(result.Id)
			util.CreateReservation(r)
		}()
	}
	return result, err
}

func (*logic) CancelReservation(id string) (model.Reservation, error) {
	result, err := _repository.Cancel(id)
	// if err == nil {
	// 	go util.CancelReservation(result)
	// }
	return result, err
}
