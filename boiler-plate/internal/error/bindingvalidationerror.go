package error

import (
	"encoding/json"
	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

// BindingValidationError ...
type BindingValidationError struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Fields  []Field `json:"fields"`
}

// Field ...
type Field struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error ...
func (e *BindingValidationError) Error() string {
	return e.Message
}

func (e *BindingValidationError) generateError(key string, messages ...string) {
	switch messages[0] {
	case "min":
		temp := &Field{
			Field:   key,
			Message: fmt.Sprintf("%s minimal adalah %s", key, messages[1]),
		}
		(*e).Fields = append((*e).Fields, *temp)

	case "max":
		temp := &Field{
			Field:   key,
			Message: fmt.Sprintf("%s maksimal adalah %s", key, messages[1]),
		}
		(*e).Fields = append((*e).Fields, *temp)
	case "len":
		temp := &Field{
			Field:   key,
			Message: fmt.Sprintf("panjang maksimal dari %s adalah %s karakter", key, messages[1]),
		}
		(*e).Fields = append((*e).Fields, *temp)
	case "required":
		temp := &Field{
			Field:   key,
			Message: fmt.Sprintf("%s harus diisi", key),
		}
		(*e).Fields = append((*e).Fields, *temp)
	}
}

// NewBindingValidationError ...
func NewBindingValidationError(ve interface{}) *BindingValidationError {
	valError := BindingValidationError{
		Code:    "ValidationError",
		Message: "validasi gagal",
	}

	if ve != nil {
		switch ve.(type) {
		case validator.ValidationErrors:
			errs := ve.(validator.ValidationErrors)
			for _, value := range errs {
				valError.generateError(value.Field, value.ActualTag, fmt.Sprintf("%v", value.Value))
			}
		case *json.UnmarshalTypeError:
			valError.Fields = append(valError.Fields, Field{
				Field:   "",
				Message: "JSON marshal error",
			})
		default:
			valError.Fields = append(valError.Fields, Field{
				Field:   "",
				Message: "unknown validation error",
			})
			break
		}

	}
	return &valError
}
