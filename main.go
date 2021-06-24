package main

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/RonnanSouza/kafka_environment/consumer"
	"github.com/RonnanSouza/kafka_environment/producer"
)

func main() {

	var (
		isConsumer = kingpin.Flag("consumer", "Use this flag to start a Kafka Consumer.").Default("false").Bool()
		isProducer = kingpin.Flag("producer", "Use this flag to start a Kafka Producer.").Default("false").Bool()
		kafkaURIs  = []string{}
	)

	kingpin.Flag("kafka.server", "Address (host:port) of Kafka server.").Required().StringsVar(&kafkaURIs)

	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	// Decision Time
	if *isProducer {
		producer.StartProducer(kafkaURIs)
	} else if *isConsumer {
		consumer.StartConsumer(kafkaURIs)
	}
}
