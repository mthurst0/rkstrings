package rkstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	assert.Equal(t, "foo_url_ref", ToSnakeCase("FooURLRef"))
	assert.Equal(t, "foo_id", ToSnakeCase("fooID"))
}
