package survey

// Survey describes a survey
// This struct represents a survey with an ID, name, list of questions, and creation timestamp
type Survey struct {
	ID        string     `json:"id" bson:"id"`                             // Unique identifier for the survey
	Name      string     `json:"name" bson:"name" validate:"required"`     // Name of the survey, required field
	Questions []Question `json:"questions" bson:"questions" validate:"required,min=2"` // List of questions, required field with minimum of 2 questions
	CreatedAt int64      `json:"createdAt" bson:"createdAt"`               // Timestamp of when the survey was created
}

// Surveys is a slice of survey pointers
// This type represents a collection of surveys
type Surveys []*Survey

// Question describes a survey question
// This struct represents a question within a survey with an ID and text
type Question struct {
	ID   int    `json:"id" bson:"id"`               // Unique identifier for the question within the survey
	Text string `json:"text" bson:"text" validate:"required"` // Text of the question, required field
}
