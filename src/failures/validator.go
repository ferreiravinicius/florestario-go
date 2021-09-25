package failures

type ValidationError struct {
	Field   string
	Message string
}

func Validation(message string, field ...string) ValidationError {
	f := append(field, "")[0]
	return ValidationError{
		Message: message,
		Field:   f,
	}
}

func (e ValidationError) Error() string {
	return e.Message
}
