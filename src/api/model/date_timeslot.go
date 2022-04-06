package model

import "time"

type DateTimeSlot struct {
	Id       string       `bson:"_id" json:"id"`
	Date     time.Time    `bson:"date" json:"date"`
	Day      time.Weekday `bson:"day" json:"day"`
	TimeSlot TimeSlot     `bson:"time_slot" json:"time_slot"`
}
