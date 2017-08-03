package metools

import (
	"io/ioutil"
	"os"
)

// CheckOrCreateDir 检查或生成路径
func CheckOrCreateDir(dir string) (string, error) {
	_, err := ioutil.ReadDir(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0777)
	}
	return dir, err
}

// SnakeName 驼峰转蛇形
func SnakeName(base string) string {
	var r = make([]rune, 0, len(base))
	var b = []rune(base)
	for i := 0; i < len(b); i++ {
		if i > 0 && b[i] >= 'A' && b[i] <= 'Z' {
			r = append(r, '_', b[i]+32)
			continue
		}
		if i == 0 && b[i] >= 'A' && b[i] <= 'Z' {
			r = append(r, b[i]+32)
			continue
		}
		r = append(r, b[i])
	}
	return string(r)
}
