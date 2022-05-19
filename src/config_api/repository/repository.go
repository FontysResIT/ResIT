package repository

import "github.com/FontysResIT/config_api/repository/mongodb"

// var Repo Repository
var DateTimeSlot IDateTimeslot
var TimeSlot ITimeSlot
var Restaurant IRestaurant

func Init() {
	mongodb.Init()
	mongoDatabase := mongodb.GetMongoDB()
	// New repositories are defined here
	DateTimeSlot = mongodb.NewDateTimeslotRepository(mongoDatabase)
	TimeSlot = mongodb.NewTimeSlotRepository(mongoDatabase)
	Restaurant = mongodb.NewRestaurantRepository(mongoDatabase)
}
