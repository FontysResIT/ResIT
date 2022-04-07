package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RealSnowKid/ResIT/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDatabase *mongo.Database

func GetMongoDB() *mongo.Database {
	return mongoDatabase
}

func Init() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	config := config.GetConfig()
	connectionString := config.GetString("mongo.url")
	fmt.Println(connectionString)
	clientOptions := options.Client().
		ApplyURI(connectionString).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	mongoDatabase = client.Database(config.GetString("mongo.database"))

}
