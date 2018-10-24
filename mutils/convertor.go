package mutils

import "encoding/json"

// 	模型转换器
type ModelConvertor interface {
	Convert(src, dest interface{}) error
}

type JsonConvert struct{}

func (j *JsonConvert) Convert(src, dest interface{}) error {
	return JSONConvert(src, dest)
}

func JSONConvert(m interface{}, dest interface{}) (err error) {
	bs, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(bs, dest)
	}
	return
}
