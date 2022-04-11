package repository

import (
	"time"

	"github.com/RealSnowKid/ResIT/model"
)

type IReservation interface {
	All() []model.Reservation
}

type ITimeSlot interface {
	All() []model.TimeSlot
}

type IDateTimeslot interface {
	All() []model.DateTimeSlot
	AllByDate(time.Time) []model.DateTimeSlot
	IdByDate(time.Time) []string
}
