package mongodb

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/FontysResIT/datetime_api/config"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDatabase *mongo.Database

func GetMongoDB() *mongo.Database {
	return mongoDatabase
}

func Init() {
	config := config.GetConfig()
	var clientOptions *options.ClientOptions
	var connectionString string
	if config.GetBool("mongo.atlas") {
		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		connectionString = fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", config.GetString("mongo.username"), url.QueryEscape(config.GetString("mongo.password")), config.GetString("mongo.url"))
		clientOptions = options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPIOptions)
	} else {
		cre := options.Credential{
			Username: config.GetString("mongo.username"),
			Password: url.QueryEscape(config.GetString("mongo.password")),
		}
		connectionString = fmt.Sprintf("mongodb://%s/", config.GetString("mongo.url"))
		clientOptions = options.Client().
			ApplyURI(connectionString).
			SetAuth(cre)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error(err)
	}
	mongoDatabase = client.Database(config.GetString("mongo.database"))

}
