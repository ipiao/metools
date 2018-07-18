package mutils

import (
	"reflect"
	"sort"
)

// Index 下标
// TODO bug
func Index(s interface{}, slice interface{}) int {
	v := reflect.ValueOf(slice)
	k := reflect.TypeOf(slice).Kind()
	if k == reflect.Slice || k == reflect.Array {
		var a = []interface{}{}
		for i := 0; i < v.Len(); i++ {
			a = append(a, v.Index(i).Interface())
		}
		return IndexA(a, func(i int) bool {
			return a[i] == s
		})
	}
	return -1
}

// IndexA 获取字符串在数组的下标，如果重复只返回第一个
func IndexA(a []interface{}, fn func(i int) bool) int {
	for j := range a {
		if fn(j) {
			return j
		}
	}
	return -1
}

// IndexS 获取字符串在数组的下标，如果重复只返回第一个
func IndexS(slice interface{}, fn func(i int) bool) int {
	v := reflect.ValueOf(slice)
	k := reflect.TypeOf(slice).Kind()
	if k == reflect.Slice || k == reflect.Array {
		var a = []interface{}{}
		for i := 0; i < v.Len(); i++ {
			a = append(a, v.Index(i).Interface())
		}
		return IndexA(a, fn)
	}
	return -1
}

// IndexStringArray 字符串在数组下标
func IndexStringArray(s string, a []string) int {
	for i := range a {
		if a[i] == s {
			return i
		}
	}
	return -1
}

func RemoveDuplicatesInts(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	base := nums[0]
	l := len(nums)
	for i := 1; i < l; {
		if nums[i] == base {
			nums = append(nums[:i], nums[i+1:]...)
			l--
		} else {
			base = nums[i]
			i++
		}
	}

	return nums
}
