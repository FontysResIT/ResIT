package repository

import "github.com/RealSnowKid/ResIT/model"

type IReservation interface {
	All() []model.Reservation
}