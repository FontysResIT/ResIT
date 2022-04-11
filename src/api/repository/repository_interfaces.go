package repository

import (
	"time"

	"github.com/RealSnowKid/ResIT/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IReservation interface {
	All() []model.Reservation
	AllByDate([]primitive.ObjectID) []model.Reservation
	Create(model.Reservation) *mongo.InsertOneResult
}

type ITimeSlot interface {
	All() []model.TimeSlot
}

type IDateTimeslot interface {
	All() []model.DateTimeSlot
	AllByDate(time.Time) []model.DateTimeSlot
	IdByDate(time.Time) []primitive.ObjectID
}
