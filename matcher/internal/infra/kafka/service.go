package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/amintasvrp/prosperity/matcher/internal/market/dto"
	"github.com/amintasvrp/prosperity/matcher/internal/market/entity"
	"github.com/amintasvrp/prosperity/matcher/internal/market/transformer"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"sync"
)

func Consume(msgChan chan *ckafka.Message, consumer *Consumer) {
	err := consumer.Consume(msgChan)
	if err != nil {
		panic(err)
	}
}

func TransformIn(msgChan chan *ckafka.Message, ordersIn chan *entity.Order, wg *sync.WaitGroup) {
	for msg := range msgChan {
		wg.Add(1)
		fmt.Println(string(msg.Value))
		tradeInput := dto.TradeInput{}
		err := json.Unmarshal(msg.Value, &tradeInput)
		if err != nil {
			panic(err)
		}
		order := transformer.TransformInput(tradeInput)
		ordersIn <- order
	}
}

func TransformOut(producer *Producer, ordersOut chan *entity.Order) {
	for res := range ordersOut {
		output := transformer.TransformOutput(res)
		outputJson, err := json.MarshalIndent(output, "", " ")
		fmt.Println(string(outputJson))
		if err != nil {
			fmt.Println(err)
		}
		err = producer.Publish(outputJson, []byte("orders"), "output")
		if err != nil {
			fmt.Println(err)
		}
	}
}
