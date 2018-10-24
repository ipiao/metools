package mutils

import (
	"bytes"
	"encoding/gob"
)

// // Cloneable 接口
// type Cloneable interface {
// 	DeepCopy(src, dest interface{}) error
// }

// // MtClone clone
// type MtClone struct{}

// DeepCopy 深复制
func DeepCopy(src, dest interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dest)
}
