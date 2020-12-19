package messagequeue

import (
	"fmt"
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/headend/share-module/configuration"
)

type MQ struct {
	Producer *kafka.Producer
	Consumer *kafka.Consumer
	Err      error
}

func (mq *MQ) PushMessage(config *configuration.Conf, msg string) *MQ {
	mq.PushMsgByTopic(config, msg, config.MQ.WarmUpTopic)
	return mq
}

func (mq *MQ) CloseProducer() *MQ {
	mq.Producer.Close()
	return mq
}

func (mq *MQ) PushMsgByTopic(config *configuration.Conf, msg string, topic string) *MQ {
	serverAddr := fmt.Sprintf("%s:%d", config.MQ.Host, config.MQ.Port)
	mq.Producer, mq.Err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": serverAddr})
	if mq.Err != nil {
		log.Print(mq.Err)
		panic(mq.Err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	// Delivery report handler for produced messages
	go func() {
		for e := range mq.Producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
					wg.Done()
					break
				} else {
					log.Printf("Delivered message to %v\n", ev.TopicPartition)
					wg.Done()
					break
				}
			}

		}
		//log.Printf("Track event closed.")

	}()

	// Produce messages to topic (asynchronously)
	_ = mq.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}, nil)
	// Wait for message deliveries
	mq.Producer.Flush(15 * 1000)
	//log.Print("wait.................................................")
	wg.Wait()
	//log.Print("Tracking action close procedure")
	mq.Producer.Close()
	return mq
}
