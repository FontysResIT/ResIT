package mongodb

import (
	"context"
	"fmt"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/FontysResIT/ResIT/config"
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
	connectionString := fmt.Sprintf("%s://%s:%s@%s?retryWrites=true&w=majority", config.GetString("mongo.prefix"), config.GetString("mongo.username"), url.QueryEscape(config.GetString("mongo.password")), config.GetString("mongo.url"))
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
