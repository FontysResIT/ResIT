package model

type GuestPersona struct {
	GuestName           string   `bson:"guest_name" json:"guest_name"`
	DietaryRequirements []string `bson:"dietary_requirements" json:"dietary_requirements"`
	Remark              string   `bson:"remark" json:"remark"`
	FoodPreferences     []string `bson:"food_preferences" json:"food_preferences"`
}
