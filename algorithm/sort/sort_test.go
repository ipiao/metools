package sort

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	is := sort.IntSlice{1, 2, 3, 23, 34, 2345, 12, 23}
	QuickSort(is)
	t.Log(is)
}
