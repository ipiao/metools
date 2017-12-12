package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// ParseFile parse file into interface
func ParseFile(path string, i interface{}, kind string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	switch kind {
	case "json":
		return json.Unmarshal(bs, i)
	case "yaml":
		return yaml.Unmarshal(bs, i)
	default:
		return fmt.Errorf("err parser kind %s", kind)
	}
}

// ParseDir 配置文件路径保持在一个文件加下，直接读取文件夹下的文件进配置文件
// i must be a struct
func ParseDir(dir string, i interface{}, kind string) error {
	v := reflect.ValueOf(i)
	return parseDir(dir, v, kind)
}

func parseDir(dir string, v reflect.Value, kind string) error {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("无法将配置文件信息解析到一个非结构体类型 %s 中！", v.Kind().String())
	}
	// get files in dir
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	// parse files
	for _, f := range fs {

		// // look up file by field name
		// fv := v.FieldByNameFunc(func(name string) bool {
		// 	return strings.TrimSuffix(f.Name(), ".yaml") == utils.SnakeName(name)
		// })

		// look up file by tag
		var fv reflect.Value
		for i := 0; i < v.NumField(); i++ {
			if tag, ok := v.Type().Field(i).Tag.Lookup(kind); ok && strings.HasSuffix(f.Name(), "."+kind) &&
				tag == strings.TrimSuffix(f.Name(), "."+kind) {
				fv = v.Field(i)
			}
		}

		if !f.IsDir() {
			if fv.CanAddr() && fv.CanInterface() {
				err = ParseFile(path.Join(dir, f.Name()), fv.Addr().Interface(), kind)
				if err != nil {
					return err
				}
			} else {
				//return fmt.Errorf("无法将文件 %s 进行配置映射", f.Name())
			}
		} else {
			if fv.CanAddr() {
				parseDir(path.Join(dir, f.Name()), fv, kind)
			}
		}
	}
	return nil
}