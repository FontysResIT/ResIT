package logic

import (
	"github.com/RealSnowKid/ResIT/model"
	"github.com/RealSnowKid/ResIT/repository"
)

type IReservationLogic interface {
	GetAllReservations() []model.Reservation
}

var _repository repository.IReservation

type logic struct {}

func NewReservationLogic(repository repository.IReservation) *logic {
	_repository = repository
	return &logic{}
}

func (* logic) GetAllReservations() []model.Reservation {
	return _repository.All()
}