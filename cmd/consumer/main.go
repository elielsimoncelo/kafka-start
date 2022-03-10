package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	//hostname, _ := os.Hostname()

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-start-kafka-1",
		"client.id":         fmt.Sprintf("goapp-consumer-%s", "1"),
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(configMap)

	if err != nil {
		fmt.Println("Erro ao consumir mensagem", err.Error())
	}

	topics := []string{"teste"}
	consumer.SubscribeTopics(topics, nil)

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			fmt.Println(string(msg.Value), string(msg.Key), msg.TopicPartition)
		}
	}
}
