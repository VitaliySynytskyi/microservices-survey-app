package repository

import (
	"context"
	"fmt"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/config"
	"github.com/jackc/pgx/v4"
	"github.com/streadway/amqp"
)

// ConnectToRabbit establishes a connection and channel to RabbitMQ
// This function returns the connection and channel, or an error if the connection fails
func ConnectToRabbit(cfg config.RabbitConfig) (*amqp.Connection, *amqp.Channel, error) {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.Username, cfg.Password, cfg.Hostname, cfg.Port)
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	err = declareQueue(ch, cfg.QueueName)
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}

// declareQueue declares a RabbitMQ queue
// This function ensures the queue is declared with the correct settings
func declareQueue(ch *amqp.Channel, queueName string) error {
	_, err := ch.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	return err
}

// ConnectToPostgres establishes a connection to the PostgreSQL database
// This function returns the connection or an error if the connection fails
func ConnectToPostgres(cfg config.PostgresConfig) (*pgx.Conn, error) {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Hostname,
		cfg.Port,
		cfg.Database,
	)
	conn, err := pgx.Connect(context.Background(), addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
