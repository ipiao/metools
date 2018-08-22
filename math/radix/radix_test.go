package radix

import (
	"testing"
)

func TestRadix(t *testing.T) {
	n := NewNumber(-19, 16)
	t.Log(n.mods)
	t.Log(n.sign)
	t.Log(n.Int())
	t.Log(n.String())
}
