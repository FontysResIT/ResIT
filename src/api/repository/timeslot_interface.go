package repository

import "github.com/RealSnowKid/ResIT/model"

type ITimeSlot interface {
	All() []model.TimeSlot
}
