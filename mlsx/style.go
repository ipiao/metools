package mlsx

import "github.com/tealeg/xlsx"

// DefaultTitleStyle 默认样式
func DefaultTitleStyle() *xlsx.Style {
	return &xlsx.Style{
		Border: xlsx.Border{
			Left:        "",
			LeftColor:   "",
			Right:       "",
			RightColor:  "",
			Top:         "",
			TopColor:    "",
			Bottom:      "",
			BottomColor: "",
		},
		ApplyBorder: true,

		Font: xlsx.Font{
			Name:   "微软雅黑",
			Size:   20,
			Family: 2,
			Bold:   true,
		},
		ApplyFont: true,

		Fill: xlsx.Fill{
			PatternType: "",
			BgColor:     "",
			FgColor:     "",
		},
		ApplyFill: false,

		Alignment: xlsx.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		ApplyAlignment: true,
	}
}

// DefaultInfoStyle 默认样式
func DefaultInfoStyle() *xlsx.Style {
	return &xlsx.Style{
		Border: xlsx.Border{
			Left:        "",
			LeftColor:   "",
			Right:       "",
			RightColor:  "",
			Top:         "",
			TopColor:    "",
			Bottom:      "",
			BottomColor: "",
		},
		ApplyBorder: true,

		Font: xlsx.Font{
			Name:   "微软雅黑",
			Size:   9,
			Family: 2,
		},
		ApplyFont: true,

		// Fill: xlsx.Fill{
		// 	PatternType: "solid",
		// 	BgColor:     "",
		// 	FgColor:     "",
		// },
		// ApplyFill: false,

		Alignment: xlsx.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		ApplyAlignment: true,
	}
}

// DefaultCellStyle 默认样式
func DefaultCellStyle() *xlsx.Style {
	return &xlsx.Style{
		Border: xlsx.Border{
			Left:        "",
			LeftColor:   "",
			Right:       "",
			RightColor:  "",
			Top:         "",
			TopColor:    "",
			Bottom:      "",
			BottomColor: "",
		},
		ApplyBorder: true,

		//	Fill: xlsx.Fill{
		//		PatternType: "solid",
		//		BgColor:     "FFCCFFFF",
		//		FgColor:     "FFDCE6F2",
		//	},
		//	ApplyFill: false,

		Font: xlsx.Font{
			Name:   "微软雅黑",
			Size:   10,
			Family: 2,
		},
		ApplyFont: true,

		Alignment: xlsx.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		ApplyAlignment: true,
	}
}
