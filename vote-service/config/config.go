package config

import (
	"time"

	"github.com/joeshaw/envdecode"
)

// Config stores complete application configuration
// This struct holds the configuration for HTTP, RabbitMQ, Survey gRPC, and Postgres
type Config struct {
	HTTP       HTTPConfig
	Rabbit     RabbitConfig
	SurveyGrpc SurveyGrpcConfig
	Postgres   PostgresConfig
}

// HTTPConfig stores HTTP configuration
// This struct holds the configuration for the HTTP server
type HTTPConfig struct {
	Hostname     string        `env:"HTTP_HOSTNAME"`                         // Hostname for the HTTP server
	Port         uint16        `env:"HTTP_PORT,default=8082"`                // Port for the HTTP server
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT,default=5s"`          // Read timeout for the HTTP server
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT,default=10s"`        // Write timeout for the HTTP server
	IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT,default=2m"`          // Idle timeout for the HTTP server
}

// RabbitConfig stores RabbitMQ configuration
// This struct holds the configuration for RabbitMQ
type RabbitConfig struct {
	Hostname  string `env:"RABBITMQ_HOSTNAME,default=localhost"`             // Hostname for RabbitMQ
	Port      uint16 `env:"RABBITMQ_PORT,default=5672"`                      // Port for RabbitMQ
	Username  string `env:"RABBITMQ_USER,default=guest"`                     // Username for RabbitMQ
	Password  string `env:"RABBITMQ_PASSWORD,default=guest"`                 // Password for RabbitMQ
	QueueName string `env:"RABBITMQ_QUEUE,default=votes"`                    // Queue name for RabbitMQ
}

// SurveyGrpcConfig stores configuration to connect to the survey gRPC service
// This struct holds the configuration for the survey gRPC service
type SurveyGrpcConfig struct {
	Hostname string `env:"SURVEY_GRPC_HOSTNAME,default=localhost"`           // Hostname for the survey gRPC service
	Port     uint16 `env:"SURVEY_GRPC_PORT,default=9000"`                    // Port for the survey gRPC service
}

// PostgresConfig stores Postgres configuration
// This struct holds the configuration for Postgres
type PostgresConfig struct {
	Hostname string `env:"POSTGRES_HOSTNAME,default=localhost"`              // Hostname for Postgres
	Port     uint16 `env:"POSTGRES_PORT,default=5432"`                       // Port for Postgres
	User     string `env:"POSTGRES_USER,default=admin"`                      // Username for Postgres
	Password string `env:"POSTGRES_PASSWORD,default=admin"`                  // Password for Postgres
	Database string `env:"POSTGRES_DB,default=voting"`                       // Database name for Postgres
	Tables   PostgresTablesConfig                                            // Tables configuration for Postgres
}

// PostgresTablesConfig stores Postgres tables configuration
// This struct holds the table configuration for Postgres
type PostgresTablesConfig struct {
	Results string `env:"POSTGRES_TABLES_RESULTS,default=results"`           // Table name for results
}

// GetConfig loads and returns application configuration
// This function loads the configuration using envdecode and returns it
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg)
	return cfg, err
}
