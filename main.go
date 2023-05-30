package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	connect, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	defer connect.Close()
	fmt.Println("successfully connected to rabbitMQ")
}
