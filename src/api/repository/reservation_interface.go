package repository

import (
	"github.com/RealSnowKid/ResIT/model"
)

type IReservation interface {
	All() []model.Reservation
	Create(model.Reservation) (model.Reservation, error)
}
