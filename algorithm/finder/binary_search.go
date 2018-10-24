package finder

// Comparator 比较器
type Comparator interface {
	Less(i, j int) bool
	Len() int
	Compare(i int, val interface{}) int
}

// IntSlice 实现比较器
type IntSlice []int

// Less 实现
func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

// Len 实现
func (is IntSlice) Len() int {
	return len(is)
}

// Compare  实现
func (is IntSlice) Compare(i int, val interface{}) int {
	n, ok := val.(int)
	if !ok {
		panic("compared value must be kind of int")
	}
	if is[i] < n {
		return -1
	} else if is[i] > n {
		return 1
	}
	return 0
}

// BinarySearch 二分查找
func BinarySearch(cmp Comparator, val interface{}) int {
	low := 0
	high := cmp.Len()
	for low < high {
		mid := (low + high) / 2
		ret := cmp.Compare(mid, val)
		if ret == -1 {
			low = mid
		} else if ret == 1 {
			high = mid
		} else {
			return mid
		}
	}
	return -1
}
