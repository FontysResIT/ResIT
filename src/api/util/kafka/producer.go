package kafka

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"time"

	"github.com/RealSnowKid/ResIT/model"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var writer *kafka.Writer
var topicPrefix string

type KafkaProducer struct{}

func NewProducer(config *viper.Viper) *KafkaProducer {
	mechanism, err := scram.Mechanism(scram.SHA256, config.GetString("kafka.username"), config.GetString("kafka.password"))
	if err != nil {
		log.Error(err)
	}
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:     config.GetStringSlice("kafka.brokers"),
		Logger:      kafka.LoggerFunc(log.Infof),
		ErrorLogger: kafka.LoggerFunc(log.Errorf),
		Dialer: &kafka.Dialer{
			Timeout:       10 * time.Second,
			DualStack:     true,
			TLS:           &tls.Config{},
			SASLMechanism: mechanism,
		},
	})
	topicPrefix = config.GetString("kafka.topic_prefix")
	return &KafkaProducer{}
}

func (*KafkaProducer) CreateReservation(reservation model.Reservation) {
	key, _ := reservation.Id.MarshalJSON()
	reservationJson, _ := json.Marshal(reservation)
	err := writer.WriteMessages(context.TODO(), kafka.Message{
		Topic: topicPrefix + "reservationapi.reservation.create",
		Key:   key,
		Value: reservationJson,
	})
	if err != nil {
		log.Error(err)
	}
}
