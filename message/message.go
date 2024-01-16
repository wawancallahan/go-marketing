package message

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"matsukana.cloud/go-marketing/config"
)

var (
	QueueNames = map[string]string{
		"go_firebase": "go.firebase",
		"go_email":    "go.email",
	}
)

type Message struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewMessage(config *config.Config) *Message {
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

	if err := channelRabbitMQ.ExchangeDeclare("go", "topic", true, false, false, false, nil); err != nil {
		panic(err)
	}

	for queueName, routingKey := range QueueNames {
		queueDeclare, err := channelRabbitMQ.QueueDeclare(
			queueName, // name
			false,     // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)

		if err != nil {
			panic(err)
		}

		if err := channelRabbitMQ.QueueBind(queueDeclare.Name, routingKey, "go", false, nil); err != nil {
			panic(err)
		}
	}

	return &Message{
		Connection: connectRabbitMQ,
		Channel:    channelRabbitMQ,
	}
}

func (m *Message) Consumer() {
	// Subscribing to QueueService1 for getting messages.

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	for queueName, _ := range QueueNames {
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

				type s struct {
					ID   int
					Name string
				}

				var result s

				err := json.Unmarshal(message.Body, &result)

				if err != nil {
					log.Println(err)
				}

				log.Println(result.ID)
			}
		}()
	}

	<-forever
}

func (m *Message) Publish(queueName string) {

}
