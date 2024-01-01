package rkstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, "FooUrlRef", ToCamelCase("foo_url_ref"))
	assert.Equal(t, "FooId", ToCamelCase("foo_id"))
}

func TestSnakeCase(t *testing.T) {
	assert.Equal(t, "foo_url_ref", ToSnakeCase("FooURLRef"))
	assert.Equal(t, "foo_id", ToSnakeCase("fooID"))
}

func TestIsLowerCase(t *testing.T) {
	assert.True(t, IsLowerCase("foo"))
	assert.False(t, IsLowerCase("Foo"))
	assert.False(t, IsLowerCase("fOo"))
}

func TestIsUpperCase(t *testing.T) {
	assert.True(t, IsUpperCase("FOO"))
	assert.False(t, IsUpperCase("Foo"))
	assert.False(t, IsUpperCase("fOO"))
}

func TestLowerCaseFirst(t *testing.T) {
	assert.Equal(t, "foo", LowerCaseFirst("Foo"))
	assert.Equal(t, "fOo", LowerCaseFirst("fOo"))
}

func TestUpperCaseFirst(t *testing.T) {
	assert.Equal(t, "Foo", UpperCaseFirst("foo"))
	assert.Equal(t, "FOo", UpperCaseFirst("fOo"))
}
