package util

import (
	"github.com/FontysResIT/config_api/config"
	"github.com/FontysResIT/config_api/model"
	"github.com/FontysResIT/config_api/util/confluent"
	log "github.com/sirupsen/logrus"
)

type IProducer interface {
	CreateRestaurant(model.Restaurant)
}

var _producers []IProducer

func InitProducers() {
	config := config.GetConfig()
	if config.GetBool("kafka.enabled") {
		log.Info("Initializing with a kafka producer")
		// _producers = append(_producers, kafka.NewProducer(config))
		_producers = append(_producers, confluent.NewProducer(config))
	}
}

func CreateRestaurant(restaurant model.Restaurant) {
	for _, producer := range _producers {
		go producer.CreateRestaurant(restaurant)
	}
}
