package validation

import "strings"

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e FieldError) Error() string {
	if e.Field == "" {
		return e.Message
	}
	return e.Field + ": " + e.Message
}

type Errors []FieldError

func (e *Errors) Add(field, message string) {
	*e = append(*e, FieldError{Field: field, Message: message})
}

func (e *Errors) AddGlobal(message string) {
	e.Add("", message)
}

func (e Errors) Any() bool {
	return len(e) > 0
}

func (e Errors) HasField(field string) bool {
	for _, fe := range e {
		if fe.Field == field {
			return true
		}
	}
	return false
}

func (e Errors) ForField(field string) []FieldError {
	var out []FieldError
	for _, fe := range e {
		if fe.Field == field {
			out = append(out, fe)
		}
	}
	return out
}

func (e Errors) First(field string) (FieldError, bool) {
	for _, fe := range e {
		if fe.Field == field {
			return fe, true
		}
	}
	return FieldError{}, false
}

func (e Errors) Error() string {
	if len(e) == 0 {
		return "validation: no errors"
	}
	parts := make([]string, 0, len(e))
	for _, fe := range e {
		parts = append(parts, fe.Error())
	}
	return strings.Join(parts, "; ")
}
