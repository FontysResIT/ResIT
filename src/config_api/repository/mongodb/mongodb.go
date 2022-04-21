package mongodb

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/FontysResIT/config_api/config"
	log "github.com/sirupsen/logrus"

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
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", config.GetString("mongo.username"), url.QueryEscape(config.GetString("mongo.password")), config.GetString("mongo.url"), config.GetString("mongo.database"))
	clientOptions := options.Client().
		ApplyURI(connectionString).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error(err)
	}
	mongoDatabase = client.Database(config.GetString("mongo.database"))

}
