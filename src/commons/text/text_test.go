package text_test

import (
	"pesthub/commons/text"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	assert.Equal(t, "iaac", text.Normalize("íãáç"))
	assert.Equal(t, "abc", text.Normalize("a@b?c!#"))
	assert.Equal(t, "", text.Normalize(""))
}

func TestSlugfy(t *testing.T) {
	assert.Equal(t, "iaac-macarrao-zzx", text.Slugfy("íãáç macarrão zzX"))
}
