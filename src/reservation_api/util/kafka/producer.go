package kafka

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"time"

	"github.com/FontysResIT/ResIT/model"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var writer *kafka.Writer
var topicPrefix string

type KafkaProducer struct{}

func NewProducer(config *viper.Viper) *KafkaProducer {
	var mechanism sasl.Mechanism
	var username = config.GetString("kafka.username")
	var password = config.GetString("kafka.password")
	switch config.GetString("kafka.mechanism") {
	case "plain":
		mechanism = plainMechanism(username, password)
		log.Infof("PLAIN %s", mechanism)
	case "sasl":
		mechanism = saslMechanism(username, password)
	default:
		mechanism = nil
	}
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:     config.GetStringSlice("kafka.brokers"),
		Logger:      kafka.LoggerFunc(log.Infof),
		ErrorLogger: kafka.LoggerFunc(log.Errorf),
		Dialer: &kafka.Dialer{
			Timeout:       10 * time.Second,
			DualStack:     true,
			TLS:           &tls.Config{InsecureSkipVerify: true},
			SASLMechanism: mechanism,
		},
	})
	topicPrefix = config.GetString("kafka.topic_prefix")
	return &KafkaProducer{}
}

func saslMechanism(username string, password string) sasl.Mechanism {
	mechanism, err := scram.Mechanism(scram.SHA256, username, password)
	if err != nil {
		log.Error(err)
	}
	return mechanism

}

func plainMechanism(username string, password string) plain.Mechanism {
	return plain.Mechanism{Username: username, Password: password}
}

func (*KafkaProducer) CreateReservation(reservation model.ReservationReadDTO) {
	key, _ := reservation.Id.MarshalJSON()
	reservationJson, _ := json.Marshal(reservation)
	log.Infoln("Producing CreateReservation")
	err := writer.WriteMessages(context.TODO(), kafka.Message{
		Topic: topicPrefix + "reservationapi.reservation.create",
		Key:   key,
		Value: reservationJson,
	})
	if err != nil {
		log.Error(err)
	}
}

func (*KafkaProducer) CancelReservation(reservation model.ReservationReadDTO) {
	key, _ := reservation.Id.MarshalJSON()
	reservationJson, _ := json.Marshal(reservation)
	err := writer.WriteMessages(context.TODO(), kafka.Message{
		Key:   key,
		Value: reservationJson,
	})
	if err != nil {
		log.Error(err)
	}
}
