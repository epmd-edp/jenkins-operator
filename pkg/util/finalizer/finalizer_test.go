package finalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsMethod(t *testing.T) {
	assert.True(t, ContainsString([]string{"a", "b"}, "b"))
}
