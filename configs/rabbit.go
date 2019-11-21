package configs

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

// this function for connect to host RabbitMQ get in ENV
func connectRabbit() *amqp.Connection {
	host := os.Getenv("AMQP_URL")
	// checking connect to host AMQ
	conn, err := amqp.Dial(host)
	if err != nil {
		log.Println(err)
	}

	return conn
}

// Publish for Channel AMQ return channel
func Publish(label string, data string) {
	conn := connectRabbit()
	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	// defer ch.Close()

	q, err := ch.QueueDeclare(
		label, // name
		true,  // durable
		false, // delete when unused
		false, // exclusif
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, //routing key
		false,  //mandatory
		false,  // immediati
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(data),
		})

	if err != nil {
		panic(err)
	}
}
