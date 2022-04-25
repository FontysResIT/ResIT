package repository

import (
	"github.com/FontysResIT/ResIT/repository/mongodb"
)

// var Repo Repository
var Reservation IReservation
var DateTimeSlot IDateTimeslot
var TimeSlot ITimeSlot

func Init() {
	mongodb.Init()
	mongoDatabase := mongodb.GetMongoDB()
	Reservation = mongodb.NewReservationRepository(mongoDatabase)
	// New repositories are defined here
	DateTimeSlot = mongodb.NewDateTimeslotRepository(mongoDatabase)
	TimeSlot = mongodb.NewTimeSlotRepository(mongoDatabase)
}
