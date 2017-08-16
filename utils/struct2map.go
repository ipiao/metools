package metools

import (
	"encoding/json"
	"reflect"
)

// Struct2Map turn struct to map
func Struct2Map(obj interface{}) map[string]interface{} {
	return struct2map(reflect.ValueOf(obj))
}

func struct2map(v reflect.Value) map[string]interface{} {
	k := v.Kind()
	if k == reflect.Ptr || k == reflect.Interface {
		v = v.Elem()
	}
	t := v.Type()
	if v.Kind() == reflect.Struct {
		var data = make(map[string]interface{})
		for i := 0; i < t.NumField(); i++ {
			if v.Field(i).CanInterface() {
				tag := t.Field(i).Tag.Get("map")
				if tag == "-" {
					continue
				} else if tag == "" {
					tag = SnakeName(t.Field(i).Name)
				}
				if (v.Field(i).Kind() == reflect.Struct || v.Field(i).Kind() == reflect.Ptr) && v.Field(i).Type().String() != "time.Time" {
					data[tag] = struct2map(v.Field(i))
				} else {
					data[tag] = v.Field(i).Interface()
				}
			}

		}
		return data
	}
	return nil
}

// StructMarshalToMap turn struct to map by using jsonmarshal
// bennchmerk 不通过
func StructMarshalToMap(obj interface{}) map[string]interface{} {
	bs, _ := json.Marshal(obj)
	var res = make(map[string]interface{})
	json.Unmarshal(bs, &res)
	return res
}
