package messagequeue

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/headend/share-module/configuration"
)

func (mq *MQ) InitConsumer(config *configuration.Conf) *MQ {
	serverAddr := fmt.Sprintf("%s:%d", config.MQ.Host, config.MQ.Port)
	mq.Consumer, mq.Err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": serverAddr,
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
	})
	if mq.Err != nil {
		panic(mq.Err)
	}
	_ = mq.Consumer.SubscribeTopics([]string{config.MQ.WarmUpTopic, "^aRegex.*[Tt]opic"}, nil)
	return mq
}

func (mq *MQ) InitConsumerByTopic(config *configuration.Conf,topic string) *MQ {
	serverAddr := fmt.Sprintf("%s:%d", config.MQ.Host, config.MQ.Port)
	mq.Consumer, mq.Err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": serverAddr,
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
	})
	if mq.Err != nil {
		panic(mq.Err)
	}
	_ = mq.Consumer.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)
	return mq
}

func (mq *MQ) CloseConsumer() *MQ {
	_ = mq.Consumer.Close()
	return mq
}
