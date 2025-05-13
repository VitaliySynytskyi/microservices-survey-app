package survey

// Serializer contains functions to encode and decode surveys
// This interface defines the methods required for serializing and deserializing survey data
type Serializer interface {
	// Encode encodes a survey with status
	// This method converts a survey with status into a byte slice (e.g., JSON)
	Encode(survey *SurveyWithStatus) ([]byte, error)

	// EncodeMultiple encodes multiple surveys with status
	// This method converts a collection of surveys with status into a byte slice (e.g., JSON)
	EncodeMultiple(surveys []SurveyWithStatus) ([]byte, error)

	// EncodeSurvey encodes a survey without status
	// This method converts a survey into a byte slice (e.g., JSON)
	EncodeSurvey(survey *Survey) ([]byte, error)

	// EncodeErrorResponse encodes an error response
	// This method converts an error response into a byte slice (e.g., JSON)
	EncodeErrorResponse(er ErrorResponse) ([]byte, error)

	// EncodeSurveyResponse encodes a survey response
	// This method converts a survey response into a byte slice (e.g., JSON)
	EncodeSurveyResponse(sr SurveyResponse) ([]byte, error)

	// EncodeSurveysResponse encodes a surveys response
	// This method converts a surveys response into a byte slice (e.g., JSON)
	EncodeSurveysResponse(sr SurveysResponse) ([]byte, error)

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
