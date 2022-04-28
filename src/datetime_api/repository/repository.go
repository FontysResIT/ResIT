package repository

import "github.com/FontysResIT/datetime_api/repository/mongodb"

// var Repo Repository
var DateTimeSlot IDateTimeslot
var TimeSlot ITimeSlot

func Init() {
	mongodb.Init()
	mongoDatabase := mongodb.GetMongoDB()
	// New repositories are defined here
	DateTimeSlot = mongodb.NewDateTimeslotRepository(mongoDatabase)
	TimeSlot = mongodb.NewTimeSlotRepository(mongoDatabase)
}
