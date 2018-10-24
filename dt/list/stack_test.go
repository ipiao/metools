package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack(10)
	assert.Equal(t, s.Len(), 0)
	assert.Nil(t, s.Top())
	assert.Nil(t, s.Pop())
	s.Push(1)
	assert.Equal(t, s.Pop(), 1)
}
