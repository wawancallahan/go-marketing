package message

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"matsukana.cloud/go-marketing/config"
)

type Message struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func New(config *config.Config) *Message {
	amqpServerURL := config.GetString("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}

	return &Message{
		Connection: connectRabbitMQ,
		Channel:    channelRabbitMQ,
	}
}

func (m *Message) Consumer() {
	// Subscribing to QueueService1 for getting messages.

	consumeQueueNames := map[string]string{}

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	for queueName, _ := range consumeQueueNames {
		messages, _ := m.Channel.Consume(
			queueName, // queue name
			"",        // consumer
			true,      // auto-ack
			false,     // exclusive
			false,     // no local
			false,     // no wait
			nil,       // arguments
		)

		go func() {
			for message := range messages {
				// For example, show received message in a console.
				log.Printf(" > Received message: %s\n", message.Body)
			}
		}()
	}

	<-forever
}
