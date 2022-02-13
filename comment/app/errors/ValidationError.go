package errors

// ValidationError Represents Validation type error
type ValidationError struct {
	Err string `json:"error" example:"Name must be specified"`
}

// Error Implements error interface
func (err ValidationError) Error() string {
	return err.Err
}

// NewValidationError returns new instance of Validation error.
func NewValidationError(error string) *ValidationError {
	return &ValidationError{
		Err: error,
	}
}
