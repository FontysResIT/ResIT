package repository

import (
	"time"

	"github.com/FontysResIT/ResIT/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IReservation interface {
	All() []model.Reservation
	AllByDate([]primitive.ObjectID) []model.Reservation
	Create(model.Reservation) (model.Reservation, error)
}

type ITimeSlot interface {
	All() []model.TimeSlot
	Create(model.TimeSlot) (model.TimeSlot, error)
}

type IDateTimeslot interface {
	All() []model.DateTimeSlot
	AllByDate(time.Time) []model.DateTimeSlot
	IdByDate(time.Time) []primitive.ObjectID
}
