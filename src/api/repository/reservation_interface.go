package repository

import (
	"github.com/RealSnowKid/ResIT/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IReservation interface {
	All() []model.Reservation
	Create(model.Reservation) *mongo.InsertOneResult
}