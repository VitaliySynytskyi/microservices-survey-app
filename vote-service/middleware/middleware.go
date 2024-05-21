package middleware

import (
	"context"
	"net/http"
	"sync"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/serializer"
	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"
)

type key int

// SerializerKey is used as a key to store the serializer in the context
const SerializerKey key = 0

var (
	// serializers stores instances of different serializers
	serializers      = make(map[string]vote.Serializer)
	serializersMutex sync.RWMutex
)

// init initializes the default serializers
func init() {
	registerSerializer("application/json", serializer.NewVoteJSONSerializer())
	// Add more serializers here if needed
}

// registerSerializer registers a serializer for a given content type
func registerSerializer(contentType string, s vote.Serializer) {
	serializersMutex.Lock()
	defer serializersMutex.Unlock()
	serializers[contentType] = s
}

// getSerializer retrieves the serializer for the given content type
func getSerializer(contentType string) vote.Serializer {
	serializersMutex.RLock()
	defer serializersMutex.RUnlock()
	if s, exists := serializers[contentType]; exists {
		return s
	}
	// Default to JSON serializer if none match
	return serializers["application/json"]
}

// AddSerializer adds a serializer to the request context
// This middleware function attaches a serializer to the context of each request
func AddSerializer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the appropriate serializer based on the content-type
		serializer := getSerializer(r.Header.Get("Content-Type"))

		// Store the serializer in context so it can be used by handlers
		ctx := context.WithValue(r.Context(), SerializerKey, serializer)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
