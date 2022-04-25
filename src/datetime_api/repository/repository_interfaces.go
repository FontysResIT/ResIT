package repository

import (
	"time"

	"github.com/FontysResIT/datetime_api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITimeSlot interface {
	All() []model.TimeSlot
	Create(model.TimeSlot) (model.TimeSlot, error)
}

type IDateTimeslot interface {
	All() []model.DateTimeSlot
	AllByDate(time.Time) []model.DateTimeSlot
	IdByDate(time.Time) []primitive.ObjectID
}
