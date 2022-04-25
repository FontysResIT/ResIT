package mongodb

import (
	"context"
	"fmt"

	"github.com/RealSnowKid/ResIT/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *MongoDBTimeSlot) Create(timeslot model.TimeSlot) (model.TimeSlot, error) {
	collection := repo.db.Collection("timeslots")
	timeslot.Id = primitive.NewObjectID()
	result, err := collection.InsertOne(context.TODO(), timeslot)
	// timeslot.Id = result.InsertedID.(primitive.ObjectID)
	log.Println("Inserted ID", result.InsertedID)
	log.Println("Timeslot ID:", timeslot.Id)
	if err != nil {
		log.Error(err)
	}
	return timeslot, err
}
