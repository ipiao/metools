package finder

import "testing"

func TestBinaryFinder(t *testing.T) {
	is := IntSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
	for i := 0; i < len(is); i++ {
		ind := BinarySearch(is, i)
		ind2 := BinarySearch(is, is[i])
		t.Log(ind, " --", ind2)
	}
}
