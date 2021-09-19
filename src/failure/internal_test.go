package failure_test

import (
	"errors"
	"pesthub/failure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalError(t *testing.T) {

	fakeErr := errors.New("fake error")

	e1 := failure.InternalError{Cause: fakeErr}
	assert.Error(t, e1)
	assert.NotNil(t, e1.Cause)
	assert.Equal(t, fakeErr.Error(), e1.Cause.Error())
	assert.NotEmpty(t, e1.Error())

	e2 := failure.Internal(fakeErr)
	assert.Error(t, e2)
	assert.NotNil(t, e2.Cause)
	assert.Equal(t, fakeErr.Error(), e2.Cause.Error())
	assert.NotEmpty(t, e2.Error())
}
