package vote

// Service contains functions to create and load votes
// This interface defines the methods required for managing votes in the application
type Service interface {
	// Insert stores a new vote
	// This method saves a new vote to the repository
	Insert(vote *Vote) error

	// GetResults gets the results for a given survey
	// This method retrieves the results for the specified survey ID
	GetResults(surveyID string) (Results, error)
}
