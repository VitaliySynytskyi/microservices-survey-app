package vote

// AnswerType defines the type of an answer
type AnswerType string

// Available answer types
const (
	AnswerTypeOption AnswerType = "option" // Option selection
	AnswerTypeText   AnswerType = "text"   // Text answer
	AnswerTypeRating AnswerType = "rating" // Rating value
	AnswerTypeScale  AnswerType = "scale"  // Scale value
	AnswerTypeDate   AnswerType = "date"   // Date answer
)

// Vote describes a vote
// This struct represents a vote with an ID, survey ID, question ID, and timestamp
type Vote struct {
	ID          string     `json:"id"`                                 // Unique identifier for the vote
	Survey      string     `json:"survey" validate:"required"`         // Survey ID associated with the vote, required field
	Question    int        `json:"question" validate:"required,min=1"` // Question ID within the survey, required field with minimum value of 1
	Timestamp   int64      `json:"timestamp"`                          // Timestamp when the vote was created
	AnswerType  AnswerType `json:"answerType" validate:"required"`     // Type of the answer
	OptionID    *int       `json:"optionId,omitempty"`                 // Option ID for single/multiple choice questions
	OptionIDs   []int      `json:"optionIds,omitempty"`                // Option IDs for multiple choice questions
	TextAnswer  *string    `json:"textAnswer,omitempty"`               // Text answer for text questions
	RatingValue *int       `json:"ratingValue,omitempty"`              // Rating value for rating questions
	ScaleValue  *int       `json:"scaleValue,omitempty"`               // Scale value for scale questions
	DateAnswer  *int64     `json:"dateAnswer,omitempty"`               // Date answer as Unix timestamp
	UserID      string     `json:"userId,omitempty"`                   // Optional user ID for non-anonymous votes
}

// Results describes the results of a survey
// This struct represents the aggregated results of a survey with the survey ID, a list of results, and the last update timestamp
type Results struct {
	Survey    string            `json:"survey"`    // Survey ID associated with the results
	Results   []QuestionResults `json:"results"`   // List of results for the survey questions
	UpdatedAt int64             `json:"updatedAt"` // Timestamp when the results were last updated
}

// QuestionResults describes the voting results of a given survey question
// This struct represents the result of votes for a specific question
type QuestionResults struct {
	Question         int                `json:"question"`                   // Question ID within the survey
	TotalVotes       int                `json:"totalVotes"`                 // Total number of votes for the question
	OptionResults    []OptionResult     `json:"optionResults,omitempty"`    // Results for each option for choice questions
	TextAnswers      []TextAnswerResult `json:"textAnswers,omitempty"`      // Text answer results for text questions
	AverageRating    *float64           `json:"averageRating,omitempty"`    // Average rating for rating questions
	RatingCounts     map[int]int        `json:"ratingCounts,omitempty"`     // Count of each rating value
	AverageScale     *float64           `json:"averageScale,omitempty"`     // Average scale for scale questions
	ScaleCounts      map[int]int        `json:"scaleCounts,omitempty"`      // Count of each scale value
	DateDistribution map[string]int     `json:"dateDistribution,omitempty"` // Distribution of date answers, grouped by day
}

// OptionResult describes the result for a specific option
type OptionResult struct {
	OptionID   int     `json:"optionId"`   // Option ID
	Count      int     `json:"count"`      // Number of votes for this option
	Percentage float64 `json:"percentage"` // Percentage of votes for this option
}

// TextAnswerResult holds text answer data for reporting
type TextAnswerResult struct {
	Answer string `json:"answer"` // Text answer
	Count  int    `json:"count"`  // How many times this answer was given
}

// Result describes a legacy result format for backward compatibility
// This struct represents the result in the original format
type Result struct {
	Question int `json:"question"` // Question ID within the survey
	Votes    int `json:"votes"`    // Number of votes for the question
}
