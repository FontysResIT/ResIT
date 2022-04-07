package mongodb

import (
	"context"
	"log"

	"github.com/RealSnowKid/ResIT/model"
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
	if err = result.All(context.TODO(), &reservations); err != nil {
		panic(err)
	}
	//fmt.Println(reservations[0].GuestNeeds[0].DietaryRequirements)
	return reservations
}
