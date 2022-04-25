package mongodb

import (
	"context"
	"fmt"

	"github.com/FontysResIT/config_api/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBTimeSlot struct {
	db *mongo.Database
}

func NewTimeSlotRepository(db *mongo.Database) *MongoDBTimeSlot {
	return &MongoDBTimeSlot{
		db: db,
	}
}

func (repo *MongoDBTimeSlot) All() []model.TimeSlot {
	var timeSlots []model.TimeSlot
	collection := repo.db.Collection("timeslots")
	result, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Error(err)
	}
	for result.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.TimeSlot
		err := result.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}

		timeSlots = append(timeSlots, elem)
	}

	return timeSlots
}
