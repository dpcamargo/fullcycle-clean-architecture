package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/dpcamargo/fullcycle-clean-architecture/internal/events"
	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order created: %v\n", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitMQ := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}
	err := h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key
		false,        // mandatory
		false,        // immediate
		msgRabbitMQ,  // msg
	)
	if err != nil {
		fmt.Printf("Error publishing message to RabbitMQ: %v\n", err)
		return
	}
}
