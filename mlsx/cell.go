package mlsx

import "github.com/tealeg/xlsx"

// Cell 自定义单元格
type Cell struct {
	Content interface{} // 内容
	Height  int         // 高度，以自身为起点，向下占据高度
	Width   int         // 宽度，以自身为起点，向右占据宽度
	Row     int         // 第几行，从1开始
	Col     int         // 第几列，从1开始
	Style   *xlsx.Style // 单元格样式
}

// NewCell 初始化单元个
func NewCell(content interface{}) Cell {
	return Cell{
		Content: content,
		Height:  1,
		Width:   1,
		Style:   DefaultCellStyle(),
	}
}

// CreateCell 创建cell
func CreateCell(content interface{}, width int64, height int64, style *xlsx.Style) Cell {
	var c = Cell{
		Content: content,
		Height:  int(height),
		Width:   int(width),
	}
	if style != nil {
		c.Style = style
	} else {
		c.Style = DefaultCellStyle()
	}
	return c
}

// CreateRow 创建一行
func CreateRow(row []interface{}) []Cell {
	var res []Cell
	for _, c := range row {
		res = append(res, NewCell(c))
	}
	return res
}
