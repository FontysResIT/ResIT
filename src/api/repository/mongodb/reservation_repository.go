package mongodb

import (
	"context"

	"github.com/RealSnowKid/ResIT/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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
	for result.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.Reservation
		err := result.Decode(&elem)
		if err != nil {
			log.Error(err)
		}

		reservations = append(reservations, elem)
	}
	// var m = &model.Reservation{Id: episodes[0].name}
	return reservations
}
