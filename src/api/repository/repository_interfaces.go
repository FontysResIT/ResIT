package repository

import (
	"time"

	"github.com/RealSnowKid/ResIT/model"
)

type IReservation interface {
	All() []model.Reservation
	AllByDate([]string) []model.Reservation
	Create(model.Reservation) (model.Reservation, error)
	Cancel(string) (model.Reservation, error)
}

type ITimeSlot interface {
	All() []model.TimeSlot
}

type IDateTimeslot interface {
	All() []model.DateTimeSlot
	AllByDate(time.Time) []model.DateTimeSlot
	IdByDate(time.Time) []string
}
