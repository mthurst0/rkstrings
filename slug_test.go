package rkstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPathSlug(t *testing.T) {
	assert.Equal(t, "h-.b", ToPathSlug("-H--.-b"))
	assert.Equal(t, "h-o", ToPathSlug(" h o"))
	assert.Equal(t, "who-is/this-and-that", ToPathSlug("Who is/ThisAndThat"))
	assert.Equal(t, "who", ToPathSlug("WHO"))
	assert.Equal(t, "foo.m4a", ToPathSlug("Foo.m4a"))
	assert.Equal(t, "h-b", ToPathSlug("h--b"))
	assert.Equal(t, "h-b", ToPathSlug("h---b"))
	assert.Equal(t, "foo2", ToPathSlug("foo2"))
}

func TestToSlug(t *testing.T) {
	assert.Equal(t, "h-b", ToSlug("-H--.-b"))
	assert.Equal(t, "h-b", ToSlug("--H--.-b"))
	assert.Equal(t, "h-b", ToSlug("h---b"))
}

func TestSlugify(t *testing.T) {
	assert.Equal(t, "h-b", Slugify("-H1-.-b9", func(r rune) bool { return false }))
	assert.Equal(t, "h1-b9", Slugify("--H1-.-b9", func(r rune) bool { return r >= '0' && r <= '9' }))
}
