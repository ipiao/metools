package metools

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// CheckOrCreateDir 检查或生成路径
func CheckOrCreateDir(dir string) (string, error) {
	_, err := ioutil.ReadDir(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0777)
	}
	return dir, err
}

func CreateFile(fname string) (*os.File, error) {
	dir := filepath.Dir(fname)
	_, err := CheckOrCreateDir(dir)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(fname)
	return f, err
}
