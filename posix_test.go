package rkstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToSafePOSIXPath(t *testing.T) {
	assert.True(t, IsPermittedInPOSIXPath('/'))
	assert.False(t, IsPermittedInPOSIXFilename('/'))
	assert.True(t, IsPermittedInPOSIXPath('.'))
	assert.True(t, IsPermittedInPOSIXPath('_'))
	assert.Equal(t, "foo-bar", ConvertToSafePOSIXFilename("foo/bar"))
	assert.Equal(t, "foo/bar", ConvertToSafePOSIXPath("foo/bar"))
	assert.Equal(t, "foo/----bar", ConvertToSafePOSIXPath("foo/*&^%bar"))
}
