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
