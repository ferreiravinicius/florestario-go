package failure

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
