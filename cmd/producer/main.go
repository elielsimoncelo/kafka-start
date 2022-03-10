package main

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)

	producer := NewKafkaProducer()

	Publish("mensagem", "teste", producer, []byte("transferencia"), deliveryChan)

	go DeliveryReport(deliveryChan)

	fmt.Println("Aguardando...")

	//producer.Flush(5000)

	b := make([]byte, 1)
	os.Stdin.Read(b)

	// e := <-deliveryChan
	// msg := e.(*kafka.Message)

	// if msg.TopicPartition.Error != nil {
	// 	fmt.Println("Erro ao enviar a mensagem")
	// 	return
	// }

	// fmt.Println("Mensagem enviada para a particao: ", msg.TopicPartition)

	// producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka-start-kafka-1",
		"client.id":           "goapp-producer",
		"delivery.timeout.ms": "0",
		"acks":                "all",  // 0 - nao preciso da resposta, 1 - lider responde que recebeu, all - lider e replicas receberam
		"enable.idempotence":  "true", // neste caso acks deve ser all
	}

	producer, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Println(err.Error())
	}

	return producer
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}

	err := producer.Produce(message, deliveryChan)

	if err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar a mensagem")
			} else {
				fmt.Println("Mensagem enviada para a particao: ", ev.TopicPartition)
			}
		}
	}
}
