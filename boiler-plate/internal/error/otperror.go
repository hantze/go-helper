package error

// OtpError ...
type OtpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error ...
func (e *OtpError) Error() string {
	return e.Message
}

// NewOtpError ...
func NewOtpError(message string) *OtpError {
	return &OtpError{
		Code:    "OtpError",
		Message: message,
	}
}
