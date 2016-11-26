package demo

import (
	"fmt"
	"testing"
	"unsafe"
)

type V struct {
	b byte
	i int32
	//m int32
	j int64
}

func (this V) GetI() {
	fmt.Printf("i=%d\n", this.i)
}

func (this V) GetJ() {
	fmt.Printf("j=%d\n", this.j)
}

func Unsafe() {
	var v = new(V)
	fmt.Printf("size of V=%d\n", unsafe.Sizeof(*v))
	//对齐值 和 偏移量
	//fmt.Printf("alignof b %d offset of b %d\n", unsafe.Alignof(v.b), unsafe.Offsetof(v.b))
	fmt.Printf("alignof i %d offset of i %d\n", unsafe.Alignof(v.i), unsafe.Offsetof(v.i))
	fmt.Printf("alignof j %d offset of j %d\n", unsafe.Alignof(v.j), unsafe.Offsetof(v.j))

	// b 和 i之间有3个字节的填充
	var i = (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(4)))
	*i = 98

	// i和j之间有4个字节的填充，但是b+j的字节数没有达到最长字节8,所以b和j在同一个最大偏移单位内
	// 如果i和j之间插入一个m 虽然m只占4个字节,但是,他要独占一个最大偏移字节
	var j = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(4) + uintptr(unsafe.Sizeof(int32(0)))))
	*j = int64(100)

	v.GetI()
	v.GetJ()
}
func TestUnsafe(t *testing.T) {
	Unsafe()
}
