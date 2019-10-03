package main

import (
	"log"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendmsg(i int, msg string, wg *sync.WaitGroup) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := msg
	log.Printf("index %d", i)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)

	failOnError(err, "Failed to publish a message")

	wg.Done()
}

func main() {
	args := os.Args[1:]
	var wg sync.WaitGroup
	for i, msg := range args {
		wg.Add(1)
		go sendmsg(i, msg, &wg)
	}
	wg.Wait()
}
