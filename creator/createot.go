package creator

import "io"

// Creator is a creator
type Creator interface {
	// 执行格式化
	format() error
	// 执行构造函数
	Exec() error
	// 设置输出
	SetOutput(io.Writer)
}
