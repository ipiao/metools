package radix

var numbers = "0123456789ABCDEF"

// Number 是进制数
type Number struct {
	bs   []byte // 字节
	base uint8  // r进制，不必支持过长的数，支持16进制以内的,默认10进制
}

// // NewNumber10 创建十进制数
// func NewNumber10(n int) *Number {

// }
