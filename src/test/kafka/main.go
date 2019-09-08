package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //确认返回
	config.Producer.Partitioner = sarama.NewRandomPartitioner //随机分区
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err: ", err)
		return
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good Test")

	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err: ", err)
		return
	}

	fmt.Printf("pid: %v, offset: %v\n", pid, offset)
}
