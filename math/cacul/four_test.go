package cacul

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	assert.True(t, IsNumber("123"))
	assert.True(t, IsNumber("123.12"))
	assert.True(t, IsNumber("-123"))
	assert.True(t, IsNumber("-123.11"))
	assert.True(t, IsNumber("-123.00"))

	assert.False(t, IsNumber("-12300..0"))
}
