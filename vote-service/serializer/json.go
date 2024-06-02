package serializer

import (
	"encoding/json"

	"github.com/VitaliySynytskyi/microservices-survey-app/vote-service/vote"
)

// voteJSONSerializer is an implementation of the vote.Serializer interface for JSON
type voteJSONSerializer struct{}

// NewVoteJSONSerializer creates a new vote JSON serializer
// This function returns a new instance of voteJSONSerializer
func NewVoteJSONSerializer() vote.Serializer {
	return &voteJSONSerializer{}
}

// Encode encodes a vote into JSON format
// This method converts a vote into a byte slice (e.g., JSON)
func (s *voteJSONSerializer) Encode(v *vote.Vote) ([]byte, error) {
	return json.Marshal(v)
}

// EncodeResults encodes vote results into JSON format
// This method converts vote results into a byte slice (e.g., JSON)
func (s *voteJSONSerializer) EncodeResults(r *vote.Results) ([]byte, error) {
	return json.Marshal(r)
}

// EncodeErrorResponse encodes an error response into JSON format
// This method converts an error response into a byte slice (e.g., JSON)
func (s *voteJSONSerializer) EncodeErrorResponse(err vote.ErrorResponse) ([]byte, error) {
	return json.Marshal(err)
}

// Decode decodes a vote from JSON format
// This method converts a byte slice into a vote (e.g., from JSON)
func (s *voteJSONSerializer) Decode(data []byte) (*vote.Vote, error) {
	v := vote.Vote{}
	err := json.Unmarshal(data, &v)
	return &v, err
}

// DecodeResults decodes vote results from JSON format
// This method converts a byte slice into vote results (e.g., from JSON)
func (s *voteJSONSerializer) DecodeResults(data []byte) (*vote.Results, error) {
	r := vote.Results{}
	err := json.Unmarshal(data, &r)
	return &r, err
}

// GetContentType returns the content type for JSON
// This method returns the MIME type of the serialized data (e.g., "application/json")
func (s *voteJSONSerializer) GetContentType() string {
	return "application/json"
}
