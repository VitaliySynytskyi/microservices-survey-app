package survey

// Service contains functions to create and fetch surveys
// This interface defines the methods required for managing surveys in the application
type Service interface {
	// Insert stores a new survey
	// This method saves a new survey to the repository
	Insert(survey *Survey) error

	// LoadByID loads a survey by ID
	// This method retrieves a survey from the repository by its unique ID
	LoadByID(id string) (*SurveyWithStatus, error)

	// Load loads all surveys
	// This method retrieves all surveys from the repository
	Load() ([]SurveyWithStatus, error)

	// LoadActive loads all active surveys that haven't expired
	// This method retrieves all active non-expired surveys
	LoadActive() ([]SurveyWithStatus, error)

	// Update updates an existing survey
	// This method updates a survey in the repository
	Update(id string, survey *Survey) error

	// ActivateSurvey activates a survey
	// This method sets a survey to active
	ActivateSurvey(id string) error

	// DeactivateSurvey deactivates a survey
	// This method sets a survey to inactive
	DeactivateSurvey(id string) error

	// SetExpirationDate sets the expiration date for a survey
	// This method sets when a survey will expire
	SetExpirationDate(id string, expiresAt int64) error

	// DeleteSurvey deletes a survey by ID
	// This method removes a survey from the repository
	DeleteSurvey(id string) error
}
