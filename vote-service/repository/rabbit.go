package repository

import (
	"fmt"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/config"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"
	"github.com/streadway/amqp"
)

type rabbitVoteRepository struct {
	cfg        config.RabbitConfig
	serializer vote.Serializer
}

// NewRabbitVoteWriterRepository creates a new RabbitMQ vote writer repository
// This function initializes and returns a new RabbitMQ vote repository instance
func NewRabbitVoteWriterRepository(cfg config.RabbitConfig, sz vote.Serializer) (vote.WriterRepository, error) {
	r := &rabbitVoteRepository{
		cfg:        cfg,
		serializer: sz,
	}
	err := r.setup()
	if err != nil {
		return r, err
	}
	return r, nil
}

// setup establishes the initial RabbitMQ connection and channel
// This function ensures the RabbitMQ connection and channel are set up correctly
func (r *rabbitVoteRepository) setup() error {
	conn, ch, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()
	return nil
}

// connect establishes a connection and channel to RabbitMQ
// This method returns the connection and channel, or an error if the connection fails
func (r *rabbitVoteRepository) connect() (*amqp.Connection, *amqp.Channel, error) {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d/", r.cfg.Username, r.cfg.Password, r.cfg.Hostname, r.cfg.Port)
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	err = r.declareQueue(ch)
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}

// declareQueue declares the RabbitMQ queue
// This method declares the queue using the configured queue name
func (r *rabbitVoteRepository) declareQueue(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare(
		r.cfg.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

// Insert adds a new vote to the RabbitMQ queue
// This method encodes the vote and publishes it to the configured RabbitMQ queue
func (r *rabbitVoteRepository) Insert(v *vote.Vote) error {
	enc, err := r.serializer.Encode(v)
	if err != nil {
		return err
	}

	conn, ch, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()

	err = r.publishVote(ch, enc)
	if err != nil {
		return err
	}

	return nil
}

// publishVote publishes the encoded vote to the RabbitMQ queue
// This method handles the publishing of the vote message to the queue
func (r *rabbitVoteRepository) publishVote(ch *amqp.Channel, enc []byte) error {
	err := ch.Publish(
		"",
		r.cfg.QueueName,
		true,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  r.serializer.GetContentType(),
			Body:         enc,
		},
	)
	return err
}
