package main

import (
	messagequeue "github.com/headend/share-module/MQ"
	"github.com/headend/share-module/configuration"
	"log"
)

func main()  {
	// load config
	var conf configuration.Conf
	conf.LoadConf()
	log.Println(conf)
	//testMQConsumer(conf)
}

func testMQConsumer(conf configuration.Conf) {
	var mq messagequeue.MQ
	mq.InitConsumerByTopic(&conf, conf.MQ.OperationTopic)
	defer mq.CloseConsumer()
	if mq.Err != nil {
		log.Print(mq.Err)
	}
	log.Printf("Listen message from %s topic\n", conf.MQ.OperationTopic)
	for {
		msg, err := mq.Consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			log.Print("Se you again!")
			break
		}
		log.Println(msg)
		log.Println(string(msg.Value))
	}
}
