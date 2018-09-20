package error

// ValidateError ...
type ValidateError struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Fields  []Field `json:"fields"`
}

// Error ...
func (e *ValidateError) Error() string {
	return e.Message
}

// NewValidateError ...
func NewValidateError(field string, message string) *ValidateError {
	validateError := &ValidateError{
		Code:    "ValidationError",
		Message: "validasi gagal",
	}
	fld := Field{
		Field:   field,
		Message: message,
	}
	validateError.Fields = append(validateError.Fields, fld)
	return validateError
}
