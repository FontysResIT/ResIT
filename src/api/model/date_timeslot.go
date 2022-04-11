package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DateTimeSlot struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Date     time.Time          `bson:"date" json:"date" binding:"Required"`
	Day      time.Weekday       `bson:"day" json:"day" binding:"Required"`
	TimeSlot TimeSlot           `bson:"time_slot" json:"time_slot" binding:"Required"`
}
