package queueHandler

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/port"
)

type (
	q struct {
		helper   port.IHelper
		consumer *kafka.Consumer
		topic    string
		done     chan struct{}
	}

	IOrderHandler interface {
		processMappingOrder(pctx context.Context, txn string)
		StartQueue()
	}
)

func (q *q) consumeMessages() {
	q.consumer.SubscribeTopics([]string{q.topic}, nil)

	for {
		select {
		case <-q.done:
			return
		default:
			if msg, err := q.consumer.ReadMessage(-1); err == nil {
				switch topicName := string(*msg.TopicPartition.Topic); topicName {
				case q.topic:
					fmt.Println("Received queue message")
					// implements
				}
			} else {
				fmt.Printf("Error while consuming message: %v (%v)\n", err, msg)
			}
		}
	}
}

func (s *q) setupShutdownSignalHandling() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Received signal. Shutting down...")
		close(s.done)
	}()
}

func (q *q) StartQueue() {
	go q.consumeMessages()
	fmt.Println("process mapping started ...")
	// Handle graceful shutdown
	q.setupShutdownSignalHandling()
	// Block until done
	<-q.done
	q.consumer.Close()
}

func (q *q) processMappingOrder(pctx context.Context, txnJson string) {
}

func NewQueueHandler(hepler port.IHelper, topic string, consumer *kafka.Consumer) IOrderHandler {
	return &q{
		helper:   hepler,
		consumer: consumer,
		topic:    topic,
		done:     make(chan struct{}),
	}
}
