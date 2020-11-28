package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/headend/share-module/configuration"
	"github.com/headend/share-module/MQ"
)

type teststruct struct {
	ID int `json:"ID"`
}

func init() {
	err := configuration.LoadConfFromFile("/tmp/testmq.yml")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {
	for {
		var mqread messagequeue.MQ
		mqread.InitConsumerByTopic(configuration.GlobalConf, "testsendread-reply1")
		log.Print("testsendread-reply:", configuration.GlobalConf)
		defer mqread.CloseConsumer()
		for {
			log.Print("ReadMessage testsendread-reply wait for:")
			msg, err := mqread.Consumer.ReadMessage(-1)
			log.Print("da nhan duoc testsendread-reply")
			wumsg := teststruct{}

			if err == nil {
				log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
				if json.Unmarshal(msg.Value, &wumsg) != nil {
					log.Print("Consumer JsonLoad error: ", err)
				}
				log.Print(wumsg)
				Jsondump, err := json.Marshal(teststruct{ID: wumsg.ID})
				checkError(err)
				var mq messagequeue.MQ
				mq.PushMsgByTopic(configuration.GlobalConf, string(Jsondump), "testsendread1")
				if mq.Err != nil {
					log.Print("error push: ", mq.Err)
				}
				break
			} else {
				log.Printf("Consumer error: %v (%v)\n", err, msg)
				break
			}
		}
	}

}
