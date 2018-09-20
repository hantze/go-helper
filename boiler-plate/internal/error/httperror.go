package error

// HTTPError ...
type HTTPError struct {
	Code     string `json:"code"`
	HttpCode int    `json:"httpCode"`
	Message  string `json:"message"`
}

// Error ...
func (e *HTTPError) Error() string {
	return e.Message
}

// NewHTTPError ...
func NewHTTPError(httpCode int) *HTTPError {
	return &HTTPError{
		Code:     "HTTPError",
		HttpCode: httpCode,
	}
}
