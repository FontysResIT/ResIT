package repository

import (
	"github.com/RealSnowKid/ResIT/repository/mongodb"
)

// var Repo Repository
var Reservation IReservation
var DateTimeSlot IDateTimeslot

func Init() {
	mongodb.Init()
	mongoDatabase := mongodb.GetMongoDB()
	Reservation = mongodb.NewReservationRepository(mongoDatabase)
	// New repositories are defined here
	DateTimeSlot = mongodb.NewDateTimeslotRepository(mongoDatabase)
	mongodb.NewTimeslotRepository(mongoDatabase)
}
