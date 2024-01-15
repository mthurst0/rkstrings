package rkstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsASCIIAlpha(t *testing.T) {
	assert.True(t, IsASCIIAlpha("abc"))
	assert.True(t, IsASCIIAlpha("ABC"))
	assert.False(t, IsASCIIAlpha("abc1"))
	assert.False(t, IsASCIIAlpha("abc!"))
	assert.False(t, IsASCIIAlpha("abc "))
	assert.False(t, IsASCIIAlpha(""))
}

func TestIsASCIIAlphaNumeric(t *testing.T) {
	assert.True(t, IsASCIIAlphaNumeric("abc1"))
	assert.True(t, IsASCIIAlphaNumeric("ABC1"))
	assert.False(t, IsASCIIAlphaNumeric("abc!"))
	assert.False(t, IsASCIIAlphaNumeric("abc "))
	assert.False(t, IsASCIIAlphaNumeric(""))
}

func TestIsASCIINumeric(t *testing.T) {
	assert.True(t, IsASCIINumeric("123"))
	assert.False(t, IsASCIINumeric("abc"))
	assert.False(t, IsASCIINumeric("123 "))
	assert.False(t, IsASCIINumeric(""))
}
