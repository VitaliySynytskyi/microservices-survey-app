package config

import (
	"os"
	"testing"
	"time"
)

// Helper function to set environment variables for testing
func setEnv(key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		panic(err)
	}
}

// Helper function to clear environment variables after testing
func clearEnv(keys ...string) {
	for _, key := range keys {
		err := os.Unsetenv(key)
		if err != nil {
			panic(err)
		}
	}
}

// TestLoadConfig tests the loading of configuration
func TestLoadConfig(t *testing.T) {
	// Set environment variables for testing
	setEnv("HTTP_HOSTNAME", "localhost")
	setEnv("HTTP_PORT", "8082")
	setEnv("HTTP_READ_TIMEOUT", "6s")
	setEnv("HTTP_WRITE_TIMEOUT", "11s")
	setEnv("HTTP_IDLE_TIMEOUT", "3m")
	setEnv("GRPC_NETWORK", "udp")
	setEnv("GRPC_HOSTNAME", "grpc.localhost")
	setEnv("GRPC_PORT", "9001")
	setEnv("MONGO_URL", "mongodb://localhost:27018")
	setEnv("MONGO_DB", "test_db")
	setEnv("MONGO_TIMEOUT", "6s")
	setEnv("REPOSITORY", "postgres")

	// Clear environment variables after testing
	defer clearEnv(
		"HTTP_HOSTNAME",
		"HTTP_PORT",
		"HTTP_READ_TIMEOUT",
		"HTTP_WRITE_TIMEOUT",
		"HTTP_IDLE_TIMEOUT",
		"GRPC_NETWORK",
		"GRPC_HOSTNAME",
		"GRPC_PORT",
		"MONGO_URL",
		"MONGO_DB",
		"MONGO_TIMEOUT",
		"REPOSITORY",
	)

	// Load the configuration
	var cfg Config
	err := loadConfig(&cfg)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Validate HTTPConfig
	if cfg.HTTP.Hostname != "localhost" {
		t.Errorf("Expected HTTP.Hostname to be 'localhost', got '%s'", cfg.HTTP.Hostname)
	}
	if cfg.HTTP.Port != 8082 {
		t.Errorf("Expected HTTP.Port to be 8082, got %d", cfg.HTTP.Port)
	}
	if cfg.HTTP.ReadTimeout != 6*time.Second {
		t.Errorf("Expected HTTP.ReadTimeout to be 6s, got %s", cfg.HTTP.ReadTimeout)
	}
	if cfg.HTTP.WriteTimeout != 11*time.Second {
		t.Errorf("Expected HTTP.WriteTimeout to be 11s, got %s", cfg.HTTP.WriteTimeout)
	}
	if cfg.HTTP.IdleTimeout != 3*time.Minute {
		t.Errorf("Expected HTTP.IdleTimeout to be 3m, got %s", cfg.HTTP.IdleTimeout)
	}

	// Validate GrpcConfig
	if cfg.Grpc.Network != "udp" {
		t.Errorf("Expected Grpc.Network to be 'udp', got '%s'", cfg.Grpc.Network)
	}
	if cfg.Grpc.Hostname != "grpc.localhost" {
		t.Errorf("Expected Grpc.Hostname to be 'grpc.localhost', got '%s'", cfg.Grpc.Hostname)
	}
	if cfg.Grpc.Port != 9001 {
		t.Errorf("Expected Grpc.Port to be 9001, got %d", cfg.Grpc.Port)
	}

	// Validate MongoConfig
	if cfg.Mongo.URL != "mongodb://localhost:27018" {
		t.Errorf("Expected Mongo.URL to be 'mongodb://localhost:27018', got '%s'", cfg.Mongo.URL)
	}
	if cfg.Mongo.DB != "test_db" {
		t.Errorf("Expected Mongo.DB to be 'test_db', got '%s'", cfg.Mongo.DB)
	}
	if cfg.Mongo.Timeout != 6*time.Second {
		t.Errorf("Expected Mongo.Timeout to be 6s, got %s", cfg.Mongo.Timeout)
	}

	// Validate Repository
	if cfg.Repository != "postgres" {
		t.Errorf("Expected Repository to be 'postgres', got '%s'", cfg.Repository)
	}
}

// TestDefaultConfig tests the loading of default configuration values
func TestDefaultConfig(t *testing.T) {
	// Clear environment variables to test default values
	clearEnv(
		"HTTP_HOSTNAME",
		"HTTP_PORT",
		"HTTP_READ_TIMEOUT",
		"HTTP_WRITE_TIMEOUT",
		"HTTP_IDLE_TIMEOUT",
		"GRPC_NETWORK",
		"GRPC_HOSTNAME",
		"GRPC_PORT",
		"MONGO_URL",
		"MONGO_DB",
		"MONGO_TIMEOUT",
		"REPOSITORY",
	)

	// Load the configuration
	var cfg Config
	err := loadConfig(&cfg)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Validate HTTPConfig with default values
	if cfg.HTTP.Port != 8081 {
		t.Errorf("Expected HTTP.Port to be 8081, got %d", cfg.HTTP.Port)
	}
	if cfg.HTTP.ReadTimeout != 5*time.Second {
		t.Errorf("Expected HTTP.ReadTimeout to be 5s, got %s", cfg.HTTP.ReadTimeout)
	}
	if cfg.HTTP.WriteTimeout != 10*time.Second {
		t.Errorf("Expected HTTP.WriteTimeout to be 10s, got %s", cfg.HTTP.WriteTimeout)
	}
	if cfg.HTTP.IdleTimeout != 2*time.Minute {
		t.Errorf("Expected HTTP.IdleTimeout to be 2m, got %s", cfg.HTTP.IdleTimeout)
	}

	// Validate GrpcConfig with default values
	if cfg.Grpc.Network != "tcp" {
		t.Errorf("Expected Grpc.Network to be 'tcp', got '%s'", cfg.Grpc.Network)
	}
	if cfg.Grpc.Port != 9000 {
		t.Errorf("Expected Grpc.Port to be 9000, got %d", cfg.Grpc.Port)
	}

	// Validate MongoConfig with default values
	if cfg.Mongo.URL != "mongodb://localhost:27017" {
		t.Errorf("Expected Mongo.URL to be 'mongodb://localhost:27017', got '%s'", cfg.Mongo.URL)
	}
	if cfg.Mongo.DB != "surveys" {
		t.Errorf("Expected Mongo.DB to be 'surveys', got '%s'", cfg.Mongo.DB)
	}
	if cfg.Mongo.Timeout != 5*time.Second {
		t.Errorf("Expected Mongo.Timeout to be 5s, got %s", cfg.Mongo.Timeout)
	}

	// Validate Repository with default value
	if cfg.Repository != "mongo" {
		t.Errorf("Expected Repository to be 'mongo', got '%s'", cfg.Repository)
	}
}
