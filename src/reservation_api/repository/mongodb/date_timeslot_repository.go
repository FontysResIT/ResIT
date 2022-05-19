package mongodb

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/FontysResIT/ResIT/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	lookup := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "timeslots"}, {Key: "localField", Value: "time_slot"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "time_slot"}}}}
	unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$time_slot"}, {Key: "preserveNullAndEmptyArrays", Value: false}}}}
	result, err := collection.Aggregate(context.TODO(), mongo.Pipeline{lookup, unwind})
	if err != nil {
		log.Error(err)
		return dateTimeSlots
	}
	if err = result.All(context.TODO(), &dateTimeSlots); err != nil {
		log.Error(err)
	}
	return dateTimeSlots
}

func (repo *MongoDBDateTimeslot) AllByDate(param time.Time) []model.DateTimeSlot {
	var dTSlots []model.DateTimeSlot
	collection := repo.db.Collection("date_timeslot")
	filter := bson.D{{Key: "date", Value: param}}
	result, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error(err)
	}
	for result.Next(context.TODO()) {

		var elem model.DateTimeSlot
		err := result.Decode(&elem)
		if err != nil {
			log.Println(err)
		}

		dTSlots = append(dTSlots, elem)
	}

	return dTSlots
}

func (repo *MongoDBDateTimeslot) IdByDate(param time.Time) []string {
	var dTSlots []string
	collection := repo.db.Collection("date_timeslot")
	filter := bson.D{{Key: "date", Value: param}}
	result, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error(err)
	}
	log.Println(result)
	for result.Next(context.TODO()) {

		var elem model.DateTimeSlot
		err := result.Decode(&elem)
		if err != nil {
			log.Println(err)
		}
		id := elem.Id.Hex()
		dTSlots = append(dTSlots, id)
	}
	return dTSlots
}

func (repo *MongoDBDateTimeslot) ById(id primitive.ObjectID) model.DateTimeSlot {
	var dts model.DateTimeSlot
	collection := repo.db.Collection("date_timeslot")
	filter := bson.D{{Key: "_id", Value: id}}
	collection.FindOne(context.TODO(), filter).Decode(&dts)
	return dts
}
