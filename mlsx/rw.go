package mlsx

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/tealeg/xlsx"
)

var (
	mutex sync.Mutex
)

// ReadExcel read excel
func ReadExcel(r io.Reader, rowLimit int) ([][][]string, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ReadExcelBinary(bs, rowLimit)
}

// ReadExcelBinary read bytes i
// rowLimit -1 is noLimit
func ReadExcelBinary(bs []byte, rowLimit int) ([][][]string, error) {
	f, err := xlsx.OpenBinaryWithRowLimit(bs, rowLimit)
	if err != nil {
		return nil, err
	}
	output := [][][]string{}
	for _, sheet := range f.Sheets {
		s := [][]string{}
		for _, row := range sheet.Rows {
			if row == nil {
				continue
			}
			r := []string{}
			for _, cell := range row.Cells {
				str, err := cell.FormattedValue()
				if err != nil {
					// Recover from strconv.NumError if the value is an empty string,
					// and insert an empty string in the output.
					if numErr, ok := err.(*strconv.NumError); ok && numErr.Num == "" {
						str = ""
					} else {
						return output, err
					}
				}
				r = append(r, strings.TrimSpace(str))
			}
			s = append(s, r)
		}
		output = append(output, s)
	}
	return output, nil
}

// CreateExcel 创建报表
// name 报表名
// infos 报表的附加信息(附加的统计信息等)
// data 单元格数据，被合并的单元格使用空格填充(占据1行1列)，至少一列
func CreateExcel(name string, infos []string, data [][]Cell, dir string) (string, error) {
	if len(data) == 0 {
		return "", errors.New("没有数据！")
	}
	// 初始化变量
	var err error
	var file *xlsx.File
	var sheet *xlsx.Sheet
	file = xlsx.NewFile()
	sheet, err = file.AddSheet(name) // 新建工作表，并设置工作表名
	if err != nil {
		return "", err
	}
	var maxHMerge = len(data[0]) - 1

	//写入数据
	// 表名行
	row := sheet.AddRow()
	row.SetHeightCM(1.0) // 设置表名行高1.0cm

	cell := row.AddCell()
	cell.SetString(name)               // 单运格内容，即表名
	cell.SetStyle(DefaultTitleStyle()) // 设置表名样式
	cell.Merge(maxHMerge, 0)

	// 信息行
	for _, info := range infos {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetString(info)
		cell.SetStyle(DefaultInfoStyle())
		cell.Merge(maxHMerge, 0)
	}

	// 数据行
	for _, rowData := range data {
		var row = sheet.AddRow()
		for _, colData := range rowData {
			var cell = row.AddCell()
			cell.SetValue(colData.Content)
			cell.SetStyle(colData.Style)
			// 水平合并
			if colData.Width != 0 {
				cell.HMerge = colData.Width - 1
			}
			// 垂直合并
			if colData.Height != 0 {
				cell.VMerge = colData.Height - 1
			}
		}
	}

	// 生成文件
	dir = filepath.Join(dir, time.Now().Format("2006-01-02"))
	_, err = CheckOrCreateDir(dir)
	if err != nil {
		log.Println("创建目录失败:", err)
		return "", err
	}
	var fileName = filepath.Join(dir, CreateFileName(name, "xlsx"))
	err = file.Save(fileName)
	if err != nil {
		log.Println("保存文件失败:", err)
		return "", err
	}
	return fileName, nil
}

// CheckOrCreateDir 检查或生成路径
func CheckOrCreateDir(dir string) (string, error) {
	_, err := ioutil.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0777)
		}
	}
	return dir, err
}

// CreateFileName 生成文件名
func CreateFileName(pname string, suffix string) string {
	node, _ := snowflake.NewNode(1)
	id := node.Generate().Int64()
	if suffix[0] != '.' {
		suffix = "." + suffix
	}
	return fmt.Sprintf("%s_%d%s", pname, id, suffix)
}
