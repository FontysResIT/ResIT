package repository

import (
	"time"

	"github.com/FontysResIT/ResIT/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IReservation interface {
	All() []model.ReservationReadDTO
	AllByDate([]string) []model.Reservation
	GetById(id primitive.ObjectID) model.ReservationReadDTO
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
	ById(primitive.ObjectID) model.DateTimeSlot
}
