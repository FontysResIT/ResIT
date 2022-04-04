package model

type Reservation struct {
	Id int `bson:"_id"`
	FirstName string `bson:"first_name"`
	LastName string `bson:"last_name"`
	Email string `bson:"email"`
	PhoneNumber string `bson:"phone_number"`
	Remark string `bson:"remark"`
	GuestCount int `bson:"guest_count"`
	GuestNeeds []GuestNeed `bson:"guest_needs"`
	IsCancelled bool `bson:"is_cancelled"`
	IsRescheduled bool `bson:"is_rescheduled"`
	DateTimeSlotId int `bson:"dts_id"`
}