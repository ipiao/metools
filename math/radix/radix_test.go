package radix

import (
	"testing"
)

func TestRadix(t *testing.T) {
	n := NewNumber('\u1111', 16)
	// t.Log(n.mods)
	// t.Log(n.sign)
	// t.Log(n.Int())
	t.Log(n.String())
	t.Log(n.ConvertTo(8))
	t.Log(n.ConvertTo(10))
	t.Log(n.ConvertTo(20))
}
