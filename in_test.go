package rkstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasAnyPrefix(t *testing.T) {
	assert.True(t, HasAnyPrefix("foobar", []string{"foo", "bar"}))
	assert.False(t, HasAnyPrefix("Foobar", []string{"foo", "bar"}))
	assert.False(t, HasAnyPrefix("Foobar", []string{}))
	assert.False(t, HasAnyPrefix("", []string{"foo", ""}))
}

func TestHasAnySuffix(t *testing.T) {
	assert.True(t, HasAnySuffix("foobar", []string{"foo", "bar"}))
	assert.False(t, HasAnySuffix("foobar", []string{"foo", "Bar"}))
	assert.False(t, HasAnySuffix("foobar", []string{}))
	assert.False(t, HasAnySuffix("", []string{"foo", ""}))
}

func TestContains(t *testing.T) {
	assert.True(t, Contains([]string{"foo", "bar"}, "foo"))
	assert.False(t, Contains([]string{"foo", "bar"}, "Foo"))
	assert.False(t, Contains([]string{"foo", "bar"}, ""))
	assert.False(t, Contains([]string{}, "foo"))
}

func TestContainsFold(t *testing.T) {
	assert.True(t, ContainsFold([]string{"foo", "bar"}, "foo"))
	assert.True(t, ContainsFold([]string{"foo", "bar"}, "Foo"))
	assert.False(t, ContainsFold([]string{"foo", "bar"}, ""))
	assert.False(t, ContainsFold([]string{}, "foo"))
}
