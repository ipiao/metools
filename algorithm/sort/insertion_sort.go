package sort

import (
	"sort"
)

func insertionSort(data sort.Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// InsertionSortInts 用于对整形数组的排序
func InsertionSortInts(a []int) {
	insertionSort(sort.IntSlice(a), 0, len(a))
}

// InsertionSortFloats 用于对浮点数组的排序
func InsertionSortFloats(a []float64) {
	insertionSort(sort.Float64Slice(a), 0, len(a))
}

// InsertionSortString 用于对字符数组的排序
func InsertionSortString(a []string) {
	insertionSort(sort.StringSlice(a), 0, len(a))
}

type lessSwap struct {
	Less func(i, j int) bool
	Swap func(i, j int)
}

// // InsertionSortSlice 用于对字符数组的排序
// func InsertionSortSlice(slice interface{}, less func(i, j int) bool) {
// 	rv := reflect.ValueOf(slice)
// 	swap := reflect.Swapper(slice)
// 	length := rv.Len()
// 	sort.Slice
// 	insertionSort(lessSwap{less, swap, len}, 0, length)
// }

// 时间复杂度是O(n^2),最坏的情况是倒序重排 n(n-1)/2
// 空间复杂度是O(1)
// 适合少量元素的排序（8个或以下）
// 标准库的排序是堆排序、快速排序和插入排序的结合使用
