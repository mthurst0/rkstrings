package rkstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFromSlice(t *testing.T) {
	r1, err := RemoveFromSlice([]string{"foo", "bar", "baz"}, 1)
	assert.NoError(t, err)
	assert.Equal(t, []string{"foo", "baz"}, r1)
	r2, err := RemoveFromSlice([]string{"foo", "bar", "baz"}, 0)
	assert.NoError(t, err)
	assert.Equal(t, []string{"baz", "bar"}, r2)
	r3, err := RemoveFromSlice([]string{"foo", "bar", "baz"}, 3)
	assert.Error(t, err)
	assert.Nil(t, r3)
	original := []string{"foo", "bar", "baz"}
	r4, err := RemoveFromSlice(original, -1)
	assert.Error(t, err)
	assert.Nil(t, r4)
	assert.Equal(t, []string{"foo", "bar", "baz"}, original)
}

func TestRemoveFromSlicePreserveOrder(t *testing.T) {
	r1, err := RemoveFromSlicePreserveOrder([]string{"foo", "bar", "baz"}, 1)
	assert.NoError(t, err)
	assert.Equal(t, []string{"foo", "baz"}, r1)
	r2, err := RemoveFromSlicePreserveOrder([]string{"foo", "bar", "baz"}, 0)
	assert.NoError(t, err)
	assert.Equal(t, []string{"bar", "baz"}, r2)
	r3, err := RemoveFromSlicePreserveOrder([]string{"foo", "bar", "baz"}, 3)
	assert.Error(t, err)
	assert.Nil(t, r3)
	original := []string{"foo", "bar", "baz"}
	r4, err := RemoveFromSlicePreserveOrder(original, -1)
	assert.Error(t, err)
	assert.Nil(t, r4)
	assert.Equal(t, []string{"foo", "bar", "baz"}, original)
}
