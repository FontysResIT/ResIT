package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TimeSlot struct {
	Id    primitive.ObjectID `bson:"_id"   json:"id"`
	From  int                `bson:"from"  json:"from" binding:"required"`
	Until int                `bson:"until" json:"until" binding:"required"`
}
