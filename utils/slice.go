package utils

import (
	"errors"
	"reflect"
	"strconv"
)

// 元素在数组中的下标，如果是-1表示不包含，要求元素和数组的类型保持一致
func Index(arr interface{}, a interface{}) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			Log("[Error]:程序抛出异常")
		}
	}()
	var err error
	var arrValue = reflect.ValueOf(arr)
	var arrKind = arrValue.Kind()
	if arrKind != reflect.Array && arrKind != reflect.Slice {
		err = errors.New("arg1的类型必须是数组或切片")
		return -1, err
	}
	if arrValue.Len() == 0 {
		err = errors.New("数组不能为空")
		return -1, err
	}
	var aValue = reflect.ValueOf(a)
	if arrValue.Index(0).Kind() != aValue.Kind() || arrValue.Index(0).Type().Name() != aValue.Type().Name() {
		err = errors.New("类型不匹配")
		return -1, err
	}
	// 直接使用interface值比较
	for i := 0; i < arrValue.Len(); i++ {
		if arrValue.Index(i).Interface() == aValue.Interface() {
			return i, nil
		}
	}
	return -1, nil
}

// 把其他类型数组转化为string数组,主要是int
func ConvertToStringSlice(arr interface{}) ([]string, error) {
	defer func() {
		if r := recover(); r != nil {
			Log("[Error]:程序抛出异常")
		}
	}()
	var res []string
	var err error
	var arrValue = reflect.ValueOf(arr)
	var arrKind = arrValue.Kind()
	if arrKind != reflect.Array && arrKind != reflect.Slice {
		err = errors.New("arg1的类型必须是数组或切片")
		return nil, err
	}
	if arrValue.Len() == 0 {
		err = errors.New("数组不能为空")
		return nil, err
	}
	var relKind = arrValue.Index(0).Kind()
	switch relKind {
	case reflect.String:
		for i := 0; i < arrValue.Len(); i++ {
			res = append(res, arrValue.Index(i).String())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		for i := 0; i < arrValue.Len(); i++ {
			var r string
			r = strconv.FormatInt(arrValue.Index(i).Int(), 10)
			res = append(res, r)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		for i := 0; i < arrValue.Len(); i++ {
			var r string
			r = strconv.FormatUint(arrValue.Index(i).Uint(), 10)
			res = append(res, r)
		}
		//	case reflect.Float32, reflect.Float64:
		//		for i := 0; i < arrValue.Len(); i++ {
		//			var r string
		//			r = strconv.FormatFloat(arrValue.Index(i).Float(), 2, 10)
		//			res = append(res, r)
		//		}
	}
	return res, nil
}
