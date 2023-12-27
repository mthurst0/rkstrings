package rkstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, "h-.b", PathSlug("-H--.-b"))
	assert.Equal(t, "h-o", PathSlug(" h o"))
	assert.Equal(t, "who-is/this-and-that", PathSlug("Who is/ThisAndThat"))
	assert.Equal(t, "who", PathSlug("WHO"))
	assert.Equal(t, "foo.m4a", PathSlug("Foo.m4a"))
	assert.Equal(t, "h-b", PathSlug("h--b"))
	assert.Equal(t, "h-b", PathSlug("h---b"))
}
