package main

import (
	"log"

	"matsukana.cloud/go-marketing/config"
	"matsukana.cloud/go-marketing/message"
)

func main() {
	config := config.New()
	message := message.NewMessage(config)

	if message != nil {
		defer message.Connection.Close()
		defer message.Channel.Close()

		message.Consumer()

		log.Printf(" > Waiting Message Consumer!")
	} else {
		log.Printf(" > Failed Run Message Consumer!")
	}
}
