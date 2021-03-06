package mutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// CheckOrCreateDir 检查或生成路径
func CheckOrCreateDir(dir string) (string, error) {
	_, err := ioutil.ReadDir(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
	}
	return dir, err
}

// CreateFile 创建文件
func CreateFile(fname string) (*os.File, error) {
	dir := filepath.Dir(fname)
	_, err := CheckOrCreateDir(dir)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(fname)
	return f, err
}
