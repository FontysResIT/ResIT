package repository

import (
	"time"

	"github.com/FontysResIT/config_api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITimeSlot interface {
	All() []model.TimeSlot
}

type IDateTimeslot interface {
	All() []model.DateTimeSlot
	AllByDate(time.Time) []model.DateTimeSlot
	IdByDate(time.Time) []primitive.ObjectID
}
