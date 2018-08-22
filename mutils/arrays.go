package mutils

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

func IntsToString(arr []int) string {
	s := ""
	for i, a := range arr {
		if i != 0 {
			s += ","
		}
		s += strconv.Itoa(a)
	}
	return s
}

//整形转换成字节
func IntToByte(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}

func ReversInts(bs []int) {
	if len(bs) <= 1 {
		return
	}
	head := bs[0]
	ReversInts(bs[1:])
	copy(bs[:len(bs)-1], bs[1:])
	bs[len(bs)-1] = head
}

func ReverseString(s string) string {
	bs := []byte(s)
	ReversBytes(bs)
	return string(bs)
}

func ReversBytes(bs []byte) {
	if len(bs) <= 1 {
		return
	}
	head := bs[0]
	ReversBytes(bs[1:])
	copy(bs[:len(bs)-1], bs[1:])
	bs[len(bs)-1] = head
}
