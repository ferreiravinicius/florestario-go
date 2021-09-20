package failures

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
