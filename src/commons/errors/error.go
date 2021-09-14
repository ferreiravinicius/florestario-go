package errors

import "fmt"

// Method 1: custom error with args and code
// Deal with translation on controller
// Pros: No need for usecase to depend on message.GetText
// Cons: Usecase will always depends on external impl
type BusinessError struct {
	Code string
	Args map[string]string
}

func (b BusinessError) Error() string {
	return b.Code
}

func Business(code string, args ...map[string]string) BusinessError {
	return BusinessError{
		Code: code,
		Args: append(args, nil)[0],
	}
}

type UnexpectedError struct {
	Cause error
}

func (u UnexpectedError) Error() string {
	msg := fmt.Sprintf("Unexpected error: %v", u.Cause.Error())
	return msg
}

func Unexpected(cause error) UnexpectedError {
	return UnexpectedError{
		Cause: cause,
	}
}

// Method 2: just message and usecase will take care of the text
// type UseCaseError struct {
// Message string
// }
