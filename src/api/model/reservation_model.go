package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reservation struct {
	Id             primitive.ObjectID `bson:"_id" json:"id"`
	FirstName      string             `bson:"first_name" json:"first_name" binding:"Required"`
	LastName       string             `bson:"last_name" json:"last_name" binding:"Required"`
	Email          string             `bson:"email" json:"email" binding:"Required"`
	PhoneNumber    string             `bson:"phone_number" json:"phone_number" binding:"Required"`
	Remark         string             `bson:"remark" json:"remark"`
	GuestCount     int                `bson:"guest_count" json:"guest_count" binding:"Required"`
	GuestPersona   []GuestPersona     `bson:"guest_persona" json:"guest_persona"`
	IsCancelled    bool               `bson:"is_cancelled" json:"is_cancelled"`
	IsRescheduled  bool               `bson:"is_rescheduled" json:"is_rescheduled"`
	DateTimeSlotId string             `bson:"dts_id" json:"dts_id" binding:"Required"`
}
