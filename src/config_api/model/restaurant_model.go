package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Restaurant struct {
	Id                   primitive.ObjectID `bson:"_id,omitempty"   json:"id"`
	Name                 string             `bson:"name"   json:"name" binding:"required"`
	Address              string             `bson:"address"   json:"address" binding:"required"`
	Capacity             int                `bson:"capacity"   json:"capacity" binding:"required"`
	StandardOpeningHours []OpeningHours     `bson:"standard_opening_hours"   json:"standard_opening_hours" binding:"required"`
}

type OpeningHours struct {
	Open  int `bson:"open"   json:"open" binding:"required"`
	Close int `bson:"close"   json:"close" binding:"required"`
}
