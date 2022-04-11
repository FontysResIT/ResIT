package mongodb

import (
	"context"
	"fmt"
	"log"

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
	for result.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.Reservation
		err := result.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}

		reservations = append(reservations, elem)
	}
	fmt.Println(reservations)
	// var m = &model.Reservation{Id: episodes[0].name}
	return reservations
}

func (repo *MongoDBReservation) AllByDate(dtsId []string) []model.Reservation {
	var reservations []model.Reservation
	collection := repo.db.Collection("reservations")
	var dtsIds = []primitive.ObjectID{}
	for _, v := range dtsId {
		j, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			panic(err)
		}
		dtsIds = append(dtsIds, j)
	}
	log.Println(dtsIds)
	filter := bson.M{"dts_id": bson.M{"$in": dtsIds}}
	result, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
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

func (repo *MongoDBReservation) Create(reservation model.Reservation) *mongo.InsertOneResult {
	collection := repo.db.Collection("reservations")
	result, err := collection.InsertOne(context.TODO(), reservation)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
