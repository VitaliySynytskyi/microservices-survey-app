package vote

// WriterRepository contains functions to write votes to a repository
// This interface defines the methods required for writing votes to a repository
type WriterRepository interface {
	// Insert stores a new vote
	// This method saves a new vote to the repository
	Insert(v *Vote) error
}

// ResultsRepository contains functions to read votes from a repository
// This interface defines the methods required for reading vote results from a repository
type ResultsRepository interface {
	// GetResults gets the results for a given survey
	// This method retrieves the results for the specified survey ID
	GetResults(surveyID string) (Results, error)
}
