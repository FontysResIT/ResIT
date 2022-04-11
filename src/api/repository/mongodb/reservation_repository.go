package mongodb

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/RealSnowKid/ResIT/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBReservation struct {
	db *mongo.Database
}

func NewReservationRepository(db *mongo.Database) *MongoDBReservation {
	return &MongoDBReservation{
		db: db,
	}
}

func (repo *MongoDBReservation) All() []model.Reservation {
	var reservations []model.Reservation
	collection := repo.db.Collection("reservations")
	result, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = result.All(context.TODO(), &reservations); err != nil {
		log.Error(err)
	}
	return reservations
}

func (repo *MongoDBReservation) Create(reservation model.Reservation) (model.Reservation, error) {
	collection := repo.db.Collection("reservations")
	result, err := collection.InsertOne(context.TODO(), reservation)
	reservation.Id = result.InsertedID.(primitive.ObjectID)
	if err != nil {
		log.Error(err)
	}
	return reservation, err
}
