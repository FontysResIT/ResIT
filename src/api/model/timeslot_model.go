package model

type TimeSlot struct {
	Id    string `bson:"_id"   json:"id"`
	From  int    `bson:"from"  json:"from"`
	Until int    `bson:"until" json:"until"`
}
