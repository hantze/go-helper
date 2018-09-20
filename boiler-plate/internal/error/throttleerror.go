package error

// ThrottleError ...
type ThrottleError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error ...
func (e *ThrottleError) Error() string {
	return e.Message
}

// NewThrottleError ...
func NewThrottleError(message string) *ThrottleError {
	return &ThrottleError{
		Code:    "ThrottleError",
		Message: message,
	}
}
