package error

// TestError ...
type TestError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error ...
func (e *TestError) Error() string {
	return e.Message
}

// NewTestError ...
func NewTestError(message string) *TestError {
	return &TestError{
		Code:    "TestError",
		Message: message,
	}
}
