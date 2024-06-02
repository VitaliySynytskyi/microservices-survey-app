package vote

// Serializer contains functions to encode and decode votes
// This interface defines the methods required for serializing and deserializing vote data
type Serializer interface {
	// Encode encodes a vote
	// This method converts a vote into a byte slice (e.g., JSON)
	Encode(v *Vote) ([]byte, error)

	// EncodeResults encodes results
	// This method converts vote results into a byte slice (e.g., JSON)
	EncodeResults(r *Results) ([]byte, error)

	// EncodeErrorResponse encodes an error response
	// This method converts an error response into a byte slice (e.g., JSON)
	EncodeErrorResponse(err ErrorResponse) ([]byte, error)

	// Decode decodes a vote
	// This method converts a byte slice into a vote (e.g., from JSON)
	Decode(data []byte) (*Vote, error)

	// DecodeResults decodes results
	// This method converts a byte slice into vote results (e.g., from JSON)
	DecodeResults(data []byte) (*Results, error)

	// GetContentType returns the content-type
	// This method returns the MIME type of the serialized data (e.g., "application/json")
	GetContentType() string
}
