package sort

import "testing"
import stdSort "sort"

//golang 标准库的排序采用快速排序
func TestStdSort(t *testing.T) {
	var a = []int{1, 2, 6, 8, 2, 5}
	stdSort.Ints(a)
	t.Log(a)
}

func TestInsertionSort(t *testing.T) {
	var a = []int{1, 2, 6, 8, 2, 5}
	InsertionSortSlice(a, func(i, j int) bool {
		a[i] < a[j]
	})
	t.Log(a)
}
