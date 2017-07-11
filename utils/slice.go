package metools

import "reflect"

// Index 下标
func Index(s interface{}, slice interface{}) int {
	v := reflect.ValueOf(slice)
	k := reflect.TypeOf(slice).Kind()
	if k == reflect.Slice || k == reflect.Array {
		var a = []interface{}{}
		for i := 0; i < v.Len(); i++ {
			a = append(a, v.Index(i).Interface())
		}
		return IndexA(s, a, func(i int) bool {
			return a[i] == s
		})
	}
	return -1
}

// IndexA 获取字符串在数组的下标，如果重复只返回第一个
func IndexA(s interface{}, a []interface{}, fn func(i int) bool) int {
	for j := range a {
		if fn(j) {
			return j
		}
	}
	return -1
}

// IndexS 获取字符串在数组的下标，如果重复只返回第一个
func IndexS(s interface{}, slice interface{}, fn func(i int) bool) int {
	v := reflect.ValueOf(slice)
	k := reflect.TypeOf(slice).Kind()
	if k == reflect.Slice || k == reflect.Array {
		var a = []interface{}{}
		for i := 0; i < v.Len(); i++ {
			a = append(a, v.Index(i).Interface())
		}
		return IndexA(s, a, fn)
	}

	return -1
}
