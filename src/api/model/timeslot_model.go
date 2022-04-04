package model

type TimeSlot struct {
	Id int `bson:"_id"`
	From int `bson:"from"`
	Until int `bson:"until"`
}