package config

import (
	"github.com/joeshaw/envdecode"
)

// Config stores complete application configuration
// This struct aggregates all configuration required for the application
type Config struct {
	Rabbit   RabbitConfig   // RabbitMQ configuration
	Postgres PostgresConfig // PostgreSQL configuration
}

// RabbitConfig stores RabbitMQ configuration
// This struct holds the configuration settings for connecting to RabbitMQ
type RabbitConfig struct {
	Hostname  string `env:"RABBITMQ_HOSTNAME,default=localhost"` // RabbitMQ server hostname
	Port      uint16 `env:"RABBITMQ_PORT,default=5672"`          // RabbitMQ server port
	User      string `env:"RABBITMQ_USER,default=guest"`         // RabbitMQ username
	Password  string `env:"RABBITMQ_PASSWORD,default=guest"`     // RabbitMQ password
	QueueName string `env:"RABBITMQ_QUEUE,default=votes"`        // RabbitMQ queue name
}

// PostgresConfig stores Postgres configuration
// This struct holds the configuration settings for connecting to PostgreSQL
type PostgresConfig struct {
	Hostname string               `env:"POSTGRES_HOSTNAME,default=localhost"` // PostgreSQL server hostname
	Port     uint16               `env:"POSTGRES_PORT,default=5432"`          // PostgreSQL server port
	User     string               `env:"POSTGRES_USER,default=admin"`         // PostgreSQL username
	Password string               `env:"POSTGRES_PASSWORD,default=admin"`     // PostgreSQL password
	Database string               `env:"POSTGRES_DB,default=voting"`          // PostgreSQL database name
	Tables   PostgresTablesConfig // PostgreSQL tables configuration
}

// PostgresTablesConfig stores Postgres tables
// This struct holds the table names for storing votes and results in PostgreSQL
type PostgresTablesConfig struct {
	Votes   string `env:"POSTGRES_TABLE_VOTES,default=votes"`      // Table name for storing votes
	Results string `env:"POSTGRES_TABLES_RESULTS,default=results"` // Table name for storing results
}

// GetConfig loads and returns application configuration
// This function loads the configuration from environment variables
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg) // Decode environment variables into the Config struct
	return cfg, err
}
