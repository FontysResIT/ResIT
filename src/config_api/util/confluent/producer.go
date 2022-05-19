package confluent

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/FontysResIT/config_api/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var writer *kafka.Producer
var topicPrefix string

type KafkaProducer struct{}

func NewProducer(config *viper.Viper) *KafkaProducer {
	var username = config.GetString("kafka.username")
	var password = config.GetString("kafka.password")
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.GetStringSlice("kafka.brokers")[0], "security.protocol": "SASL_SSL", "sasl.mechanism": "PLAIN", "sasl.username": username, "sasl.password": password})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	topicPrefix = config.GetString("kafka.topic_prefix")
	writer = p

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	return &KafkaProducer{}
}

func (*KafkaProducer) CreateRestaurant(restaurant model.Restaurant) {
	key, _ := restaurant.Id.MarshalJSON()
	//topic := "restaurantapi.restaurant.create"
	topic := topicPrefix + "mockTopic"
	RestaurantJson, _ := json.Marshal(restaurant)
	log.Infoln("Producing CreateRestaurant")
	writer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Key:            key,
		Value:          RestaurantJson,
	}, nil)
}
