package scylla

import "github.com/RealSnowKid/ResIT/model"

type ScyllaReservation struct {
	nextId int
}

func NewReservationRepository() *ScyllaReservation {
	return &ScyllaReservation{
		nextId: 1,
	}
}

func (repo *ScyllaReservation) All() []model.Reservation {
	var reservations []model.Reservation
	var m = &model.Reservation{Id: repo.nextId}
	return append(reservations, *m )
}
