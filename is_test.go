package rkstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUUID(t *testing.T) {
	assert.True(t, IsUUID("123e4567-e89b-12d3-a456-426655440000"))
	assert.True(t, IsUUID("c73bcdcc-2669-4bf6-81d3-e4ae73fb11fd"))
	assert.True(t, IsUUID("C73BCDCC-2669-4Bf6-81d3-E4AE73FB11FD"))
	assert.False(t, IsUUID("c73bcdcc-2669-4bf6-81d3-e4an73fb11fd"))
	assert.False(t, IsUUID("c73bcdcc26694bf681d3e4ae73fb11fd"))
	assert.False(t, IsUUID("definitely-not-a-uuid"))
}
