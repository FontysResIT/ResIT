package logic

import (
	"fmt"

	"github.com/RealSnowKid/ResIT/model"
	"github.com/RealSnowKid/ResIT/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IReservationLogic interface {
	GetAllReservations() []model.Reservation
	CreateReservation(model.Reservation) *mongo.InsertOneResult
}

var _repository repository.IReservation

type logic struct{}

func NewReservationLogic(repository repository.IReservation) *logic {
	_repository = repository
	return &logic{}
}

func (*logic) GetAllReservations() []model.Reservation {
	return _repository.All()
}

func (*logic) CreateReservation(reservation model.Reservation) *mongo.InsertOneResult{
	fmt.Println("reservation logic inside createReservation")
	reservation.Id = primitive.NewObjectID()
	return _repository.Create(reservation)
}
