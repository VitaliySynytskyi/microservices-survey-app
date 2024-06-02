package queue

import (
	"fmt"
	"os"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-worker-service/config"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
)

type rabbitVoteQueue struct {
	config     config.RabbitConfig
	serializer vote.Serializer
	log        *zerolog.Logger
}

// NewRabbitVoteQueue creates a new rabbit vote queue
// This function initializes and returns a new instance of rabbitVoteQueue
func NewRabbitVoteQueue(cfg config.RabbitConfig, sz vote.Serializer, l *zerolog.Logger) VoteQueue {
	return &rabbitVoteQueue{
		config:     cfg,
		serializer: sz,
		log:        l,
	}
}

// Consume consumes votes from a RabbitMQ queue and passes them through a channel to be processed
// This method establishes a connection to RabbitMQ, declares the queue, and starts consuming messages
func (r *rabbitVoteQueue) Consume(vc chan<- *vote.Vote) {
	// Connect to RabbitMQ
	conn, ch, err := r.connectToRabbit()
	if err != nil {
		r.log.Fatal().Err(err).Msg("Cannot connect to queue")
		os.Exit(1)
	}
	defer conn.Close()
	defer ch.Close()

	// Declare the queue
	if err := r.declareQueue(ch); err != nil {
		r.log.Fatal().Err(err).Msg("Cannot declare queue")
		os.Exit(1)
	}

	// Start consuming messages
	msgs, err := ch.Consume(
		r.config.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.log.Fatal().Err(err).Msg("Unable to consume queue")
		os.Exit(1)
	}

	r.log.Info().Str("on", fmt.Sprintf("%s:%d", r.config.Hostname, r.config.Port)).Msg("Connected to queue. Awaiting messages.")

	// Process messages from the queue
	for msg := range msgs {
		r.processMessage(msg, vc)
	}
}

// connectToRabbit establishes a connection and channel to RabbitMQ
// This method returns the connection and channel, or an error if the connection fails
func (r *rabbitVoteQueue) connectToRabbit() (*amqp.Connection, *amqp.Channel, error) {
	host := fmt.Sprintf("%s:%d", r.config.Hostname, r.config.Port)
	addr := fmt.Sprintf("amqp://%s:%s@%s/", r.config.User, r.config.Password, host)
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}

// declareQueue declares the RabbitMQ queue
// This method ensures the queue is declared with the correct settings
func (r *rabbitVoteQueue) declareQueue(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare(
		r.config.QueueName,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	return err
}

// processMessage processes a single message from the queue
// This method decodes the message and sends the vote to the provided channel
func (r *rabbitVoteQueue) processMessage(msg amqp.Delivery, vc chan<- *vote.Vote) {
	v, err := r.serializer.Decode(msg.Body)
	if err != nil {
		r.log.Error().Err(err).Str("body", string(msg.Body)).Msg("Unable to parse vote from queue message")
		return
	}
	r.log.Info().Str("id", v.ID).Msg("Vote received from queue")
	vc <- v
}
