package failures

import "strings"

type InternalError struct {
	Cause error
}

func (ie InternalError) Error() string {
	return "Internal error: " + ie.Cause.Error()
}

func Internal(cause error) InternalError {
	return InternalError{Cause: cause}
}

type UseCaseError struct {
	Message string
}

func (uce UseCaseError) Error() string {
	return uce.Message
}

func UseCase(message string) UseCaseError {
	return UseCaseError{
		Message: message,
	}
}

type ErrorList []error

func (el ErrorList) Error() string {
	var sb strings.Builder
	for _, err := range el {
		sb.WriteString(err.Error())
		sb.WriteRune('\n')
	}
	return sb.String()
}
