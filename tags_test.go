package rkstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {
	tags1 := NewEmptyTags()
	assert.Equal(t, 0, tags1.Count())
	tags1.Add(Tag{
		Key:   "key1",
		Value: "value1",
	})
	assert.Equal(t, 1, tags1.Count())
}

func TestNewTags(t *testing.T) {
	tags1 := NewTags([]Tag{
		{Key: "key1", Value: "value1"},
		{Key: "key2", Value: "value2"},
	})
	assert.Equal(t, 2, tags1.Count())
	assert.Equal(t, "value1", tags1.GetWithDefault("key1", "other1"))
	assert.Equal(t, "value2", tags1.GetWithDefault("key2", "other2"))
	assert.Equal(t, "other3", tags1.GetWithDefault("key3", "other3"))
	assert.True(t, tags1.Has("key1"))
	assert.True(t, tags1.Has("key2"))
	assert.False(t, tags1.Has("key3"))
	assert.False(t, tags1.Has(""))
	tags1.Remove("key3")
	assert.Equal(t, 2, tags1.Count())
	tags1.Remove("key2")
	assert.Equal(t, 1, tags1.Count())
	assert.True(t, tags1.Has("key1"))
	assert.False(t, tags1.Has("key2"))
}
