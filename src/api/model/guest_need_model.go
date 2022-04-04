package model

type GuestNeed struct {
	GuestName string `bson:"guest_name"`
	DietaryRequirements []string `bson:"dietary_requirements"`
	Remark string `bson:"remark"`
	FoodPreferences [] string `bson:"food_preferences"`
}