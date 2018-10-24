package mdecoder

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/schema"
	"github.com/ipiao/metools/datetool"
)

// schema 将http GET url路径转换成对应的结构体

var (
	ParseTime = datetool.ParseTime
)

// SchemaDecoder url 解析
type SchemaDecoder struct {
	*schema.Decoder
}

func NewSchemaDecoder() *SchemaDecoder {
	decoder := schema.NewDecoder()
	// decoder.IgnoreUnknownKeys(true)
	// decoder.ZeroEmpty(true)
	// decoder.SetAliasTag("json")
	decoder.RegisterConverter(time.Time{}, timeConverter)
	decoder.RegisterConverter([]string{}, stringArrayConverter)
	decoder.RegisterConverter([]int64{}, int64ArrayConverter)
	decoder.RegisterConverter([]int32{}, int32ArrayConverter)
	decoder.RegisterConverter([]int8{}, int8ArrayConverter)
	decoder.RegisterConverter([]int16{}, int16ArrayConverter)
	decoder.RegisterConverter([]int{}, intArrayConverter)
	return &SchemaDecoder{decoder}
}

// 自定义修改时间解析器
func timeConverter(s string) reflect.Value {
	if v, err := ParseTime(s); err == nil {
		return reflect.ValueOf(v)
	}
	return reflect.Value{}
}

func stringArrayConverter(s string) reflect.Value {
	if len(s) > 0 {
		var ss = []string{}
		var err error
		if s[0] == '[' {
			err = json.Unmarshal([]byte(s), &ss)
		} else if s[0] == '"' {
			ss = append(ss, s[1:len(s)-1])
		} else {
			ss = append(ss, s)
		}
		if err == nil {
			return reflect.ValueOf(ss)
		}
	}
	return reflect.Value{}
}

func int64ArrayConverter(s string) reflect.Value {
	if len(s) > 0 {
		var ss = []int64{}
		var err error
		if s[0] == '[' {
			err = json.Unmarshal([]byte(s), &ss)
		} else if s[0] == '"' {
			n, _ := strconv.ParseInt(s[1:len(s)-1], 10, 64)
			ss = append(ss, n)
		} else {
			n, _ := strconv.ParseInt(s, 10, 64)
			ss = append(ss, n)
		}
		if err == nil {
			return reflect.ValueOf(ss)
		}
	}
	return reflect.Value{}
}

func int32ArrayConverter(s string) reflect.Value {
	if len(s) > 0 {
		var ss = []int32{}
		var err error
		if s[0] == '[' {
			err = json.Unmarshal([]byte(s), &ss)
		} else if s[0] == '"' {
			n, _ := strconv.ParseInt(s[1:len(s)-1], 10, 32)
			ss = append(ss, int32(n))
		} else {
			n, _ := strconv.ParseInt(s, 10, 32)
			ss = append(ss, int32(n))
		}
		if err == nil {
			return reflect.ValueOf(ss)
		}
	}
	return reflect.Value{}
}

func int16ArrayConverter(s string) reflect.Value {
	if len(s) > 0 {
		var ss = []int16{}
		var err error
		if s[0] == '[' {
			err = json.Unmarshal([]byte(s), &ss)
		} else if s[0] == '"' {
			n, _ := strconv.ParseInt(s[1:len(s)-1], 10, 64)
			ss = append(ss, int16(n))
		} else {
			n, _ := strconv.ParseInt(s, 10, 64)
			ss = append(ss, int16(n))
		}
		if err == nil {
			return reflect.ValueOf(ss)
		}
	}
	return reflect.Value{}
}

func int8ArrayConverter(s string) reflect.Value {
	if len(s) > 0 {
		var ss = []int8{}
		var err error
		if s[0] == '[' {
			err = json.Unmarshal([]byte(s), &ss)
		} else if s[0] == '"' {
			n, _ := strconv.ParseInt(s[1:len(s)-1], 10, 64)
			ss = append(ss, int8(n))
		} else {
			n, _ := strconv.ParseInt(s, 10, 64)
			ss = append(ss, int8(n))
		}
		if err == nil {
			return reflect.ValueOf(ss)
		}
	}
	return reflect.Value{}
}

func intArrayConverter(s string) reflect.Value {
	if len(s) > 0 {
		var ss = []int{}
		var err error
		if s[0] == '[' {
			err = json.Unmarshal([]byte(s), &ss)
		} else if s[0] == '"' {
			n, _ := strconv.Atoi(s[1 : len(s)-1])
			ss = append(ss, n)
		} else {
			n, _ := strconv.Atoi(s)
			ss = append(ss, n)
		}
		if err == nil {
			return reflect.ValueOf(ss)
		}
	}
	return reflect.Value{}
}
