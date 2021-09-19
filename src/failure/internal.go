package failure

type InternalError struct {
	Cause error
}

func (ie InternalError) Error() string {
	return "Internal error: " + ie.Cause.Error()
}

func Internal(cause error) InternalError {
	return InternalError{Cause: cause}
}
