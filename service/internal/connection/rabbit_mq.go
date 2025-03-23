package connection

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func connectRabbitmq() error {
	var err error
	rabbitmq, err = amqp091.Dial(conn.Rabbitmq)
	if err != nil {
		log.Println("error rabbitmq connect: ", err)
		rabbitmq.Close()
	}

	log.Println("rabbitmq conntected")
	return err
}
