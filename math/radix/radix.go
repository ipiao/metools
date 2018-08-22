package radix

import "github.com/ipiao/metools/mutils"

var tables = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func SetTables(s string) {
	tables = s
}

// Number 是进制数
type Number struct {
	mods []int // 字节
	base uint8 // r进制，不必支持过长的数，支持16进制以内的,默认10进制
	sign int8
}

// NewNumber10 创建十进制数
func NewNumber(num int, base uint8) *Number {
	if num < 0 {
		ret := NewNumber(-num, base)
		ret.sign = -1
		return ret
	}

	if int(base) > len(tables) {
		panic("unsupport base,please use SetTables first")
	}

	ret := &Number{
		mods: mutils.IntTInts(num, base),
		base: base,
		sign: 1,
	}

	if len(ret.mods) == 0 {
		ret.mods = append(ret.mods, 0)
	}
	return ret
}

func (n *Number) String() string {
	var bs = make([]byte, 0)
	if n.sign == -1 {
		bs = append(bs, '-')
	}
	for _, mod := range n.mods {
		bs = append(bs, tables[mod])
	}
	return string(bs)
}

// 这里必然是转换为10进制
func (n *Number) Int() int {
	return mutils.IntsTInt(n.mods, n.base) * int(n.sign)
}

func (n *Number) ConvertTo(base uint8) *Number {
	return NewNumber(n.Int(), base)
}
