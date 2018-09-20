package error

// AuthorizationError ...
type AuthorizationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error ...
func (e *AuthorizationError) Error() string {
	return e.Message
}

// NewAuthorizationError ...
func NewAuthorizationError() *AuthorizationError {
	return &AuthorizationError{
		Code:    "Error",
		Message: "Unauthorized",
	}
}
