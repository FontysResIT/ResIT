package logic

import (
	"github.com/RealSnowKid/ResIT/model"
	"github.com/RealSnowKid/ResIT/repository"
	"github.com/RealSnowKid/ResIT/util"
)

type IReservationLogic interface {
	GetAllReservations() []model.Reservation
	CreateReservation(model.Reservation) (model.Reservation, error)
}

var _repository repository.IReservation

type logic struct{}

func NewReservationLogic(repository repository.IReservation) *logic {
	_repository = repository
	return &logic{}
}

func (*logic) GetAllReservations() []model.Reservation {
	return _repository.All()
}

func (*logic) CreateReservation(reservation model.Reservation) (model.Reservation, error) {
	result, err := _repository.Create(reservation)
	if err == nil {
		go util.CreateReservation(result)
	}
	return result, err

}
