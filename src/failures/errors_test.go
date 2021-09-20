package failures_test

import (
	"errors"
	"pesthub/failures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseError(t *testing.T) {

	e1 := failures.UseCaseError{Message: "message"}
	assert.Error(t, e1)
	assert.Equal(t, "message", e1.Message)

	e2 := failures.UseCase("message")
	assert.Error(t, e2)
	assert.Equal(t, "message", e2.Message)
	assert.NotEmpty(t, e2.Error())
}

func TestInternalError(t *testing.T) {

	fakeErr := errors.New("fake error")

	e1 := failures.InternalError{Cause: fakeErr}
	assert.Error(t, e1)
	assert.NotNil(t, e1.Cause)
	assert.Equal(t, fakeErr.Error(), e1.Cause.Error())
	assert.NotEmpty(t, e1.Error())

	e2 := failures.Internal(fakeErr)
	assert.Error(t, e2)
	assert.NotNil(t, e2.Cause)
	assert.Equal(t, fakeErr.Error(), e2.Cause.Error())
	assert.NotEmpty(t, e2.Error())
}
