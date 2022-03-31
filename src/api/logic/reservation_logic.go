package logic

import (
	"github.com/RealSnowKid/ResIT/model"
	"github.com/RealSnowKid/ResIT/repository"
)

func GetAllReservations() []*model.Reservation {
	return repository.Reservation.All()
}