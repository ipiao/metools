package mutils

import (
	"bytes"
	"encoding/json"
)

// DeJSON decode json to interface
func DeJSON(data []byte, v interface{}) error {
	var decode = json.NewDecoder(bytes.NewBuffer(data))
	decode.UseNumber()
	return decode.Decode(&v)
}

// EnJSON 解析成json
func EnJSON(v interface{}) ([]byte, error) {
	var bs []byte
	bf := bytes.NewBuffer(bs)
	var encode = json.NewEncoder(bf)
	err := encode.Encode(v)
	return bf.Bytes(), err
}

// EnJSONStr 解析成json str
func EnJSONStr(v interface{}) (string, error) {
	bs, err := EnJSON(v)
	return string(bs), err
}
