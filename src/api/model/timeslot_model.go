package model

type TimeSlot struct {
	Id string `bson:"_id"`
	From int `bson:"from"`
	Until int `bson:"until"`
}