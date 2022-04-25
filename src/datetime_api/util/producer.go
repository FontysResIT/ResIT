package util

import (
	"github.com/FontysResIT/datetime_api/config"
	"github.com/FontysResIT/datetime_api/util/kafka"
	log "github.com/sirupsen/logrus"
)

type IProducer interface {
}

var _producers []IProducer

func InitProducers() {
	config := config.GetConfig()
	if config.GetBool("kafka.enabled") {
		log.Info("Initializing with a kafka producer")
		_producers = append(_producers, kafka.NewProducer(config))
	}
}
