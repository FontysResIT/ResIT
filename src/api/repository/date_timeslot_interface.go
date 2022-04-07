package repository

import "github.com/RealSnowKid/ResIT/model"

type IDateTimeslot interface {
	All() []model.DateTimeSlot
}
