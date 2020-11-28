package messagequeue

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/headend/share-module/configuration"
	"strconv"
)

func (mq *MQ) InitConsumer(config *configuration.Conf) *MQ {
	mq.Consumer, mq.Err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.MQ.Host + ":" +  strconv.Itoa(config.MQ.Port),
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
	mq.Consumer, mq.Err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.MQ.Host + ":" +  strconv.Itoa(config.MQ.Port),
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
