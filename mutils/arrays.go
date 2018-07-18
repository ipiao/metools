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
