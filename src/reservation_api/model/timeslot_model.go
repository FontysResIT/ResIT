package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TimeSlot struct {
	Id          primitive.ObjectID `bson:"_id"   json:"id"`
	FromHour    int                `bson:"from_hour"  json:"from_hour"`
	FromMinutes int                `bson:"from_minutes"  json:"from_minutes"`
	ToHour      int                `bson:"to_hour"  json:"to_hour"`
	ToMinutes   int                `bson:"to_minutes"  json:"to_minutes"`
}
