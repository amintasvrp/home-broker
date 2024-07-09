package trade

import (
	"github.com/amintasvrp/prosperity/trader/internal/infra/kafka"
	"github.com/amintasvrp/prosperity/trader/internal/market/entity"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"sync"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	msgChan := make(chan *ckafka.Message)
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "trader",
		"auto.offset.reset": "earliest",
	}
	producer := kafka.NewProducer(configMap)
	consumer := kafka.NewConsumer(configMap, []string{"input"})
	book := entity.NewBook(ordersIn, ordersOut, wg)

	go kafka.Consume(msgChan, consumer)
	go book.Trade()
	go kafka.Transform(msgChan, ordersIn, producer, ordersOut, wg)
}
