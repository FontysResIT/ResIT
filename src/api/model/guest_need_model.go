package model

type GuestNeed struct {
	Id int `bson:"_id"`
	GuestName string `bson:"guest_name"`
	DietaryRequirements []string `bson:"dietary_requirements"`
	Remark string `bson:"remark"`
	FoodPreferences [] string `bson:"food_preferences"`
}