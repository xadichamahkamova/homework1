package producer

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	Conn    *amqp.Connection
	Queue   amqp.Queue
	Channel *amqp.Channel
}

func DialProducer() (*Producer, error) {

	url := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(url)
	if err != nil {
		conn.Close()
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"task",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &Producer{
		Conn:    conn,
		Queue:   q,
		Channel: ch,
	}, nil
}
