package survey

// Serializer contains functions to encode and decode surveys
// This interface defines the methods required for serializing and deserializing survey data
type Serializer interface {
	// Encode encodes a survey
	// This method converts a survey into a byte slice (e.g., JSON)
	Encode(survey *Survey) ([]byte, error)

	// EncodeMultiple encodes multiple surveys
	// This method converts a collection of surveys into a byte slice (e.g., JSON)
	EncodeMultiple(surveys *Surveys) ([]byte, error)

	// EncodeErrorResponse encodes an error response
	// This method converts an error response into a byte slice (e.g., JSON)
	EncodeErrorResponse(er ErrorResponse) ([]byte, error)

	// Decode decodes a survey
	// This method converts a byte slice into a survey (e.g., from JSON)
	Decode(data []byte) (*Survey, error)

	// DecodeMultiple decodes multiple surveys
	// This method converts a byte slice into a collection of surveys (e.g., from JSON)
	DecodeMultiple(data []byte) (*Surveys, error)

	// GetContentType returns the content-type
	// This method returns the MIME type of the serialized data (e.g., "application/json")
	GetContentType() string
}
