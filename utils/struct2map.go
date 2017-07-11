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

// StructMarshaalToMap turn struct to map by using jsonmarshal
func StructMarshaalToMap(obj interface{}) map[string]interface{} {
	bs, _ := json.Marshal(obj)
	var res = make(map[string]interface{})
	json.Unmarshal(bs, &res)
	return res
}
