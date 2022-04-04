package model

import "time"

type DateTimeSlot struct {
	Id string `bson:"_id"`
	Date time.Time `bson:"date"`
	Day time.Weekday `bson:"day"`
	TimeSlot TimeSlot `bson:"time_slot"`
}