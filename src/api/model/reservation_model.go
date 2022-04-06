package model

type Reservation struct {
	Id             string      `bson:"_id" json:"id"`
	FirstName      string      `bson:"first_name" json:"first_name"`
	LastName       string      `bson:"last_name" json:"last_name"`
	Email          string      `bson:"email" json:"email"`
	PhoneNumber    string      `bson:"phone_number" json:"phone_number"`
	Remark         string      `bson:"remark" json:"remark"`
	GuestCount     int         `bson:"guest_count" json:"guest_count"`
	GuestNeeds     []GuestNeed `bson:"guest_needs" json:"guest_needs"`
	IsCancelled    bool        `bson:"is_cancelled" json:"is_cancelled"`
	IsRescheduled  bool        `bson:"is_rescheduled" json:"is_rescheduled"`
	DateTimeSlotId string      `bson:"dts_id" json:"dts_id"`
}
