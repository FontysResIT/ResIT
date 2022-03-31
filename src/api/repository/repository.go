package repository

import (
	"github.com/RealSnowKid/ResIT/repository/mongodb"
)
// var Repo Repository
var Reservation IReservation

func Init() {
	mongodb.Init()
	mongoDatabase := mongodb.GetMongoDB()
	Reservation = mongodb.NewReservationRepository(mongoDatabase)
		// New repositories are defined here
}




