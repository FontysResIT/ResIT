package mongodb

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/RealSnowKid/ResIT/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		log.Error(err)
	}
	for result.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.Reservation
		err := result.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}

		reservations = append(reservations, elem)
	}
	return reservations
}

func (repo *MongoDBReservation) AllByDate(dtsId []string) []model.Reservation {
	var reservations []model.Reservation
	collection := repo.db.Collection("reservations")
	log.Println("DTSID LENGHT", len(dtsId))
	var filter = bson.M{}
	if len(dtsId) < 1 {
		filter = bson.M{"dts_id": 0}
	} else {
		filter = bson.M{"dts_id": bson.M{"$in": dtsId}}
	}
	result, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error(err)
	}
	for result.Next(context.TODO()) {
		var elem model.Reservation
		err := result.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}

		reservations = append(reservations, elem)
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

func (repo *MongoDBReservation) Cancel(id string) (model.Reservation, error) {
	var err error

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
	}

	collection := repo.db.Collection("reservations")

	result := collection.FindOneAndUpdate(context.TODO(),
		bson.M{"_id": objID}, bson.D{
			{"$set", bson.D{{"is_cancelled", true}}},
		}, options.MergeFindOneAndUpdateOptions(&options.FindOneAndUpdateOptions{ReturnDocument: options.FindOneAndUpdate().ReturnDocument}))

	log.Println(result)
	if result.Err() != nil {
		log.Error(result.Err().Error())
	}
	var res model.Reservation
	err = result.Decode(&res)
	return res, err
}
