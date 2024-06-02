package config

import (
	"time"

	"github.com/joeshaw/envdecode"
)

// Config stores complete application configuration
type Config struct {
	HTTP       HTTPConfig
	Grpc       GrpcConfig
	Mongo      MongoConfig
	Repository string `env:"REPOSITORY,default=mongo"`
}

// HTTPConfig stores HTTP configuration
// Code Smell: Duplicated Code
// Refactoring: Using composition to reduce duplication
type HTTPConfig struct {
	Hostname     string        `env:"HTTP_HOSTNAME"`					// Hostname for the HTTP server
	Port         uint16        `env:"HTTP_PORT,default=8081"`			// Port for the HTTP server
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT,default=5s"`		// Read timeout for the HTTP server
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT,default=10s"`	// Write timeout for the HTTP server
	IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT,default=2m"`		// Idle timeout for the HTTP server
}

// GrpcConfig stores gRPC configuration
type GrpcConfig struct {
	Network  string `env:"GRPC_NETWORK,default=tcp"`
	Hostname string `env:"GRPC_HOSTNAME"`
	Port     uint16 `env:"GRPC_PORT,default=9000"`
}

// MongoConfig stores Mongo DB configuration
type MongoConfig struct {
	URL     string        `env:"MONGO_URL,default=mongodb://localhost:27017"`
	DB      string        `env:"MONGO_DB,default=surveys"`
	Timeout time.Duration `env:"MONGO_TIMEOUT,default=5s"`
}

// GetConfig loads and returns application configuration
// Code Smell: Long Method
// Refactoring: Break method into smaller, more manageable functions
func GetConfig() (Config, error) {
	var cfg Config
	err := loadConfig(&cfg)
	return cfg, err
}

// loadConfig loads the configuration using envdecode
func loadConfig(cfg *Config) error {
	return envdecode.StrictDecode(cfg)
}
