package vote

// Vote describes a vote
// This struct represents a vote with an ID, survey ID, question ID, and timestamp
type Vote struct {
	ID        string `json:"id"`                             // Unique identifier for the vote
	Survey    string `json:"survey" validate:"required"`     // Survey ID associated with the vote, required field
	Question  int    `json:"question" validate:"required,min=1"` // Question ID within the survey, required field with minimum value of 1
	Timestamp int64  `json:"timestamp"`                      // Timestamp when the vote was created
}

// Results describes the results of a survey
// This struct represents the aggregated results of a survey with the survey ID, a list of results, and the last update timestamp
type Results struct {
	Survey    string   `json:"survey"`   // Survey ID associated with the results
	Results   []Result `json:"results"`  // List of results for the survey questions
	UpdatedAt int64    `json:"updatedAt"`// Timestamp when the results were last updated
}

// Result describes the voting results of a given survey question
// This struct represents the result of a vote for a specific question with the question ID and the number of votes
type Result struct {
	Question int `json:"question"` // Question ID within the survey
	Votes    int `json:"votes"`    // Number of votes for the question
}
