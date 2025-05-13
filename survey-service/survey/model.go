package survey

import "time"

// QuestionType defines the type of a question
type QuestionType string

// Available question types
const (
	QuestionTypeSingleChoice   QuestionType = "single_choice"   // Single choice from multiple options
	QuestionTypeMultipleChoice QuestionType = "multiple_choice" // Multiple choices from options
	QuestionTypeText           QuestionType = "text"            // Open text answer
	QuestionTypeRating         QuestionType = "rating"          // Rating on a scale
	QuestionTypeScale          QuestionType = "scale"           // Likert scale or other numeric scale
	QuestionTypeDate           QuestionType = "date"            // Date selection
)

// MediaType defines the type of media attached to a question
type MediaType string

// Available media types
const (
	MediaTypeImage MediaType = "image" // Image attachment
	MediaTypeVideo MediaType = "video" // Video attachment
	MediaTypeAudio MediaType = "audio" // Audio attachment
)

// ConditionalLogicType defines the type of conditional logic
type ConditionalLogicType string

// Available conditional logic types
const (
	LogicTypeShow ConditionalLogicType = "show" // Show question when condition is met
	LogicTypeSkip ConditionalLogicType = "skip" // Skip question when condition is met
)

// Survey describes a survey
// This struct represents a survey with an ID, name, list of questions, and creation timestamp
type Survey struct {
	ID              string     `json:"id" bson:"id"`                                               // Unique identifier for the survey
	Name            string     `json:"name" bson:"name" validate:"required"`                       // Name of the survey, required field
	Description     string     `json:"description" bson:"description"`                             // Description of the survey
	Questions       []Question `json:"questions" bson:"questions" validate:"required,min=1"`       // List of questions, required field with minimum of 1 question
	CreatedAt       int64      `json:"createdAt" bson:"createdAt"`                                 // Timestamp of when the survey was created
	ExpiresAt       int64      `json:"expiresAt,omitempty" bson:"expiresAt,omitempty"`             // Optional expiration timestamp
	Active          bool       `json:"active" bson:"active"`                                       // Whether the survey is active
	AllowAnonymous  bool       `json:"allowAnonymous" bson:"allowAnonymous"`                       // Whether anonymous responses are allowed
	ThankYouMessage string     `json:"thankYouMessage,omitempty" bson:"thankYouMessage,omitempty"` // Message to show after completion
}

// Surveys is a slice of survey pointers
// This type represents a collection of surveys
type Surveys []*Survey

// Question describes a survey question
// This struct represents a question within a survey with an ID, text, and additional properties
type Question struct {
	ID               int               `json:"id" bson:"id"`                                                 // Unique identifier for the question within the survey
	Text             string            `json:"text" bson:"text" validate:"required"`                         // Text of the question, required field
	Type             QuestionType      `json:"type" bson:"type" validate:"required"`                         // Type of the question, required field
	Required         bool              `json:"required" bson:"required"`                                     // Whether an answer is required
	Options          []Option          `json:"options,omitempty" bson:"options,omitempty"`                   // Options for choice-based questions
	Media            *Media            `json:"media,omitempty" bson:"media,omitempty"`                       // Optional media attachment
	MinValue         *int              `json:"minValue,omitempty" bson:"minValue,omitempty"`                 // Min value for scale questions
	MaxValue         *int              `json:"maxValue,omitempty" bson:"maxValue,omitempty"`                 // Max value for scale questions
	ConditionalLogic *ConditionalLogic `json:"conditionalLogic,omitempty" bson:"conditionalLogic,omitempty"` // Conditional logic for question visibility
	Placeholder      string            `json:"placeholder,omitempty" bson:"placeholder,omitempty"`           // Placeholder text for text questions
	HelpText         string            `json:"helpText,omitempty" bson:"helpText,omitempty"`                 // Additional help text for the question
}

// Option represents an answer option for choice-based questions
type Option struct {
	ID    int    `json:"id" bson:"id"`                           // Unique identifier for the option
	Text  string `json:"text" bson:"text"`                       // Text of the option
	Image string `json:"image,omitempty" bson:"image,omitempty"` // Optional image URL for the option
}

// Media represents a media attachment for a question
type Media struct {
	Type    MediaType `json:"type" bson:"type"`                           // Type of the media
	URL     string    `json:"url" bson:"url"`                             // URL of the media
	Caption string    `json:"caption,omitempty" bson:"caption,omitempty"` // Optional caption for the media
}

// ConditionalLogic represents conditional display logic for a question
type ConditionalLogic struct {
	Type             ConditionalLogicType `json:"type" bson:"type"`                                         // Type of conditional logic
	SourceQuestionID int                  `json:"sourceQuestionId" bson:"sourceQuestionId"`                 // ID of the question that this condition depends on
	SourceOptionID   int                  `json:"sourceOptionId,omitempty" bson:"sourceOptionId,omitempty"` // ID of the option that triggers this condition (for choice questions)
	SourceValue      string               `json:"sourceValue,omitempty" bson:"sourceValue,omitempty"`       // Value that triggers this condition (for other question types)
	Operator         string               `json:"operator,omitempty" bson:"operator,omitempty"`             // Operator for comparison (equals, not equals, greater than, etc.)
}

// SurveyWithStatus adds expiration status to a survey
type SurveyWithStatus struct {
	*Survey
	Status SurveyStatus `json:"status"` // Current status of the survey
}

// SurveyStatus represents the current status of a survey
type SurveyStatus string

// Available survey statuses
const (
	SurveyStatusActive    SurveyStatus = "active"    // Survey is active
	SurveyStatusInactive  SurveyStatus = "inactive"  // Survey is inactive
	SurveyStatusExpired   SurveyStatus = "expired"   // Survey has expired
	SurveyStatusScheduled SurveyStatus = "scheduled" // Survey is scheduled for future
)

// IsExpired checks if the survey has expired
func (s *Survey) IsExpired() bool {
	if s.ExpiresAt == 0 {
		return false
	}
	return time.Now().UTC().Unix() > s.ExpiresAt
}

// GetStatus returns the current status of the survey
func (s *Survey) GetStatus() SurveyStatus {
	if s.IsExpired() {
		return SurveyStatusExpired
	}
	if s.Active {
		return SurveyStatusActive
	}
	return SurveyStatusInactive
}
