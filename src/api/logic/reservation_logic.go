package logic

import (
	"log"
	"time"

	"github.com/RealSnowKid/ResIT/model"
	"github.com/RealSnowKid/ResIT/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IReservationLogic interface {
	GetAllReservations() []model.Reservation
	GetAllReservationsByDate(time.Time) []model.Reservation
	CreateReservation(model.Reservation) *mongo.InsertOneResult
}

var _repository repository.IReservation
var _DTSrepository repository.IDateTimeslot

type logic struct{}

func NewReservationLogic(repository repository.IReservation, dtsRepository repository.IDateTimeslot) *logic {
	_repository = repository
	_DTSrepository = dtsRepository
	return &logic{}
}

func (*logic) GetAllReservations() []model.Reservation {
	return _repository.All()
}

func (*logic) GetAllReservationsByDate(date time.Time) []model.Reservation {
	return _repository.AllByDate(_DTSrepository.IdByDate(date))
}

func (*logic) CreateReservation(reservation model.Reservation) *mongo.InsertOneResult {
	log.Println("reservation logic inside createReservation")
	reservation.Id = primitive.NewObjectID()
	return _repository.Create(reservation)
}
