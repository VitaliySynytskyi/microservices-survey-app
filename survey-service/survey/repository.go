package survey

// Repository contains functions to store and fetch surveys from a repository
// This interface defines the methods required for interacting with a survey repository
type Repository interface {
	// Insert stores a new survey
	// This method saves a new survey to the repository
	Insert(survey *Survey) error

	// LoadByID loads a survey by ID
	// This method retrieves a survey from the repository by its unique ID
	LoadByID(id string) (*Survey, error)

	// Load loads all surveys
	// This method retrieves all surveys from the repository
	Load() (*Surveys, error)

	// Update updates an existing survey
	// This method updates a survey in the repository
	Update(survey *Survey) error

	// Delete deletes a survey by ID
	// This method deletes a survey from the repository
	Delete(id string) error
}
