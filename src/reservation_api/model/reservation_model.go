package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reservation struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName      string             `bson:"first_name" json:"first_name" binding:"required"`
	LastName       string             `bson:"last_name" json:"last_name" binding:"required"`
	Email          string             `bson:"email" json:"email" binding:"required"`
	PhoneNumber    string             `bson:"phone_number" json:"phone_number" binding:"required"`
	Remark         string             `bson:"remark" json:"remark"`
	GuestCount     int                `bson:"guest_count" json:"guest_count" binding:"required"`
	GuestPersona   []GuestPersona     `bson:"guest_persona" json:"guest_persona"`
	IsCancelled    bool               `bson:"is_cancelled" json:"is_cancelled"`
	IsRescheduled  bool               `bson:"is_rescheduled" json:"is_rescheduled"`
	DateTimeSlotId primitive.ObjectID `bson:"dts_id" json:"dts_id" binding:"required"`
	RestaurantId   primitive.ObjectID `bson:"restaurant_id" json:"restaurant_id" binding:"required"`
}

type ReservationReadDTO struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName     string             `bson:"first_name" json:"first_name" binding:"required"`
	LastName      string             `bson:"last_name" json:"last_name" binding:"required"`
	Email         string             `bson:"email" json:"email" binding:"required"`
	PhoneNumber   string             `bson:"phone_number" json:"phone_number" binding:"required"`
	Remark        string             `bson:"remark" json:"remark"`
	GuestCount    int                `bson:"guest_count" json:"guest_count" binding:"required"`
	GuestPersona  []GuestPersona     `bson:"guest_persona" json:"guest_persona"`
	IsCancelled   bool               `bson:"is_cancelled" json:"is_cancelled"`
	IsRescheduled bool               `bson:"is_rescheduled" json:"is_rescheduled"`
	DateTimeSlot  DateTimeSlot       `bson:"datetime_slot" json:"datetime_slot" binding:"required"`
	RestaurantId  primitive.ObjectID `bson:"restaurant_id" json:"restaurant_id" binding:"required"`
}
