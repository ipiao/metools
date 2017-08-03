package metools

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
)

// Struct2Map turn struct to map
func Struct2Map(obj interface{}) map[string]interface{} {

	v := reflect.Indirect(reflect.ValueOf(obj))
	t := v.Type()
	if t.Kind() != reflect.Struct {
		panic("err kind :" + t.Kind().String())
	}
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("map")
		if tag == "-" {
			continue
		} else if tag == "" {
			tag = strings.ToLower(regexp.MustCompile(`\B[A-Z]`).ReplaceAllString(t.Field(i).Name, "_$0"))
		}
		data[tag] = v.Field(i).Interface()
	}
	return data

}

// func struct2map(v reflect.Value) map[string]interface{} {
// 	v = reflect.Indirect(v)
// 	t := v.Type()
// 	if v.Kind() == reflect.Struct {
// 		var data = make(map[string]interface{})
// 		for i := 0; i < t.NumField(); i++ {
// 			tag := t.Field(i).Tag.Get("map")
// 			if tag == "-" {
// 				continue
// 			} else if tag == "" {
// 				tag = strings.ToLower(regexp.MustCompile(`\B[A-Z]`).ReplaceAllString(t.Field(i).Name, "_$0"))
// 			}
// 		}
// 	}
// }

// StructMarshaalToMap turn struct to map by using jsonmarshal
func StructMarshaalToMap(obj interface{}) map[string]interface{} {
	bs, _ := json.Marshal(obj)
	var res = make(map[string]interface{})
	json.Unmarshal(bs, &res)
	return res
}
