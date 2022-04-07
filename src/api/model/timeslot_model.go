package model

type TimeSlot struct {
	Id          string `bson:"_id"   json:"id"`
	FromHour    int    `bson:"from_hour"  json:"from_hour"`
	FromMinutes int    `bson:"from_minutes"  json:"from_miunutes"`
	ToHour      int    `bson:"to_hour"  json:"to_hour"`
	ToMinutes   int    `bson:"to_minutes"  json:"to_miunutes"`
}
