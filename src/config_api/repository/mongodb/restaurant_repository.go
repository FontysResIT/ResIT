package mongodb

import (
	"context"
	"fmt"

	"github.com/FontysResIT/config_api/model"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRestaurant struct {
	db *mongo.Database
}

func NewRestaurantRepository(db *mongo.Database) *MongoDBRestaurant {
	return &MongoDBRestaurant{
		db: db,
	}
}

func (repo *MongoDBRestaurant) All() []model.Restaurant {
	var Restaurants []model.Restaurant
	collection := repo.db.Collection("restaurants")
	result, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Error(err)
	}
	for result.Next(context.TODO()) {
		var elem model.Restaurant
		err := result.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}

		Restaurants = append(Restaurants, elem)
	}

	return Restaurants
}

func (repo *MongoDBRestaurant) Create(restaurant model.Restaurant) (model.Restaurant, error) {
	collection := repo.db.Collection("restaurants")
	result, err := collection.InsertOne(context.TODO(), restaurant)
	if err != nil {
		log.Error(err)
	}
	restaurant.Id = result.InsertedID.(primitive.ObjectID)
	return restaurant, err
}
