package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("start")
	connect, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("failed init broker connection")
		panic(err)
	}

	ch, err := connect.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	// dead letter queue - https://aws.amazon.com/what-is/dead-letter-queue/
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(q)

	if err != nil {
		fmt.Println(err)
	}
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("successfully connected to rabbitMQ")

}
