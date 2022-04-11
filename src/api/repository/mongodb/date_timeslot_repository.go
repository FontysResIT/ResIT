package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RealSnowKid/ResIT/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBDateTimeslot struct {
	db *mongo.Database
}

func NewDateTimeslotRepository(db *mongo.Database) *MongoDBDateTimeslot {
	return &MongoDBDateTimeslot{
		db: db,
	}
}

func (repo *MongoDBDateTimeslot) All() []model.DateTimeSlot {
	var dateTimeSlots []model.DateTimeSlot
	collection := repo.db.Collection("date_timeslot")
	result, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for result.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem model.DateTimeSlot
		err := result.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}

		dateTimeSlots = append(dateTimeSlots, elem)
	}

	return dateTimeSlots
}

func (repo *MongoDBDateTimeslot) Date(date time.Time) []model.DateTimeSlot {
	var dTSlots []model.DateTimeSlot
	log.Println(date)
	collection := repo.db.Collection("date_timeslot")
	filter := bson.D{{"date", date}}
	result, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	for result.Next(context.TODO()) {

		var elem model.DateTimeSlot
		err := result.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}

		dTSlots = append(dTSlots, elem)
	}

	return dTSlots
}
