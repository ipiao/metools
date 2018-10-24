package mutils

import (
	"errors"
	"strconv"
)

// IntfaceToString 将interface转换为string
func IntfaceToString(x interface{}) (string, error) {
	switch x.(type) {
	case string:
		return x.(string), nil
	case int:
		return strconv.Itoa(x.(int)), nil
	case int32:
		return strconv.Itoa(x.(int)), nil
	case int64:
		return strconv.Itoa(x.(int)), nil
	case uint32:
		return strconv.Itoa(int(x.(uint32))), nil
	case float32:
		return strconv.FormatFloat(float64(x.(float32)), 'f', 2, 64), nil
	case float64:
		return strconv.FormatFloat(x.(float64), 'f', 2, 64), nil
	default:
		return "", errors.New("not support")
	}
}

// IntfaceToInt64 将interface转换为int64
func IntfaceToInt64(x interface{}) (int64, error) {
	switch x.(type) {
	case string:
		return strconv.ParseInt(x.(string), 10, 64)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return x.(int64), nil
	case float32, float64:
		return int64(x.(float64)), nil
	default:
		return 0, errors.New("not support")
	}
}
