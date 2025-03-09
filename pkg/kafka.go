package pkg

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pstpmn/my-golang-hexagonal-template/conf"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/port"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/utils"
)

type k struct {
	Producer *kafka.Producer
}

// Consumer implements port.IMessageQueue.
func (k *k) Consumer(topic string) error {
	panic("unimplemented")
}

// Produce implements port.IMessageQueue.
func (k *k) Produce(topic, message string) error {
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}
	if err := k.Producer.Produce(kafkaMessage, nil); err != nil {
		fmt.Printf("Failed to produce message: %v\n", err)
		return err
	}
	// Flush outstanding messages
	// k.Producer.Flush(15 * 1000) // Timeout in milliseconds
	return nil
}

func NewKafkaProducer(conf conf.Kafka) (port.IMessageQueue, error) {
	cert, err := utils.Base64Decode(conf.Base64Cert)
	if err != nil {
		return nil, err
	}

	config := &kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", conf.BrokerHost[0], conf.PORT), // kafka broker addr
		"sasl.mechanism":    conf.Mechanism,
		"security.protocol": conf.Protocol,
		"sasl.username":     conf.User,
		"sasl.password":     conf.Pass,
		"ssl.ca.pem":        cert,
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, err
	}
	return &k{Producer: producer}, nil
}

func ConnectConsumer(conf conf.Kafka) (*kafka.Consumer, error) {
	cert, err := utils.Base64Decode(conf.Base64Cert)
	if err != nil {
		return nil, err
	}

	config := &kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", conf.BrokerHost[0], conf.PORT), // kafka broker addr
		"group.id":          "console-consumer-21564",                            // consumer group
		"auto.offset.reset": "earliest",
		"sasl.mechanism":    conf.Mechanism,
		"security.protocol": conf.Protocol,
		"sasl.username":     conf.User,
		"sasl.password":     conf.Pass,
		"ssl.ca.pem":        cert,
	}
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
