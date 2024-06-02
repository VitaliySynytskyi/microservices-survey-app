package survey

// Service contains functions to create and fetch surveys
// This interface defines the methods required for managing surveys in the application
type Service interface {
	// Insert stores a new survey
	// This method saves a new survey to the repository
	Insert(survey *Survey) error

	// LoadByID loads a survey by ID
	// This method retrieves a survey from the repository by its unique ID
	LoadByID(id string) (*Survey, error)

	// Load loads all surveys
	// This method retrieves all surveys from the repository
	Load() (*Surveys, error)
}
