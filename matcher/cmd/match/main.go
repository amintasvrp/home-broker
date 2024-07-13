package main

import (
	"github.com/amintasvrp/prosperity/matcher/config"
	"github.com/amintasvrp/prosperity/matcher/internal/infra/kafka"
	"github.com/amintasvrp/prosperity/matcher/internal/market/entity"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"sync"
)

func main() {
	env := config.NewConfigEnv()
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	msgChan := make(chan *ckafka.Message)
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": env.BootstrapServers,
		"group.id":          env.GroupID,
		"auto.offset.reset": env.Offset,
	}
	producer := kafka.NewProducer(configMap)
	consumer := kafka.NewConsumer(configMap, []string{"input"})
	book := entity.NewBook(ordersIn, ordersOut, wg)

	go kafka.Consume(msgChan, consumer)
	go book.Trade()
	go kafka.TransformIn(msgChan, ordersIn, wg)
	kafka.TransformOut(producer, ordersOut)
}
