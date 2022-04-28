package repository

import (
	"time"

	"github.com/FontysResIT/ResIT/model"
)

type IReservation interface {
	All() []model.Reservation
	AllByDate([]string) []model.Reservation
	Create(model.Reservation) (model.Reservation, error)
	Cancel(string) (model.Reservation, error)
}

type ITimeSlot interface {
	All() []model.TimeSlot
	Create(model.TimeSlot) (model.TimeSlot, error)
}

type IDateTimeslot interface {
	All() []model.DateTimeSlot
	AllByDate(time.Time) []model.DateTimeSlot
	IdByDate(time.Time) []string
}
