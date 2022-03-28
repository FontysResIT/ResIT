package repository

import (
	"github.com/RealSnowKid/ResIT/repository/scylla"
)
var Repo Repository

type Repository struct{
	Reservation IReservation
	// Repositories go here
	// ...
}

func Init() {
	Repo = Repository{ 
		Reservation: scylla.NewReservationRepository(),
		// New repositories are defined here
	}
}




