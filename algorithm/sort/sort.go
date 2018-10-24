package sort // BuddleSort 冒泡排序

// https://www.cnblogs.com/eniac12/p/5329396.html

// Swap 交换
type Swap interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

// IntSlice for sort []int
type IntSlice []int

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { temp := s[i]; s[i] = s[j]; s[j] = temp }
func (s IntSlice) Len() int           { return len(s) }

// BubbleSortInts 对[]int冒泡排序
func BubbleSortInts(data []int) {
	swap := IntSlice(data)
	bubbleSort(swap)
}

// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(n^2)
// 最优时间复杂度 ---- 如果能在内部循环第一次运行时,使用一个旗标来表示有无需要交换的可能,可以把最优时间复杂度降低到O(n)
// 平均时间复杂度 ---- O(n^2)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 稳定
func bubbleSort(data Swap) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

// CocktailSortInts 对[]int鸡尾酒排序
func CocktailSortInts(data []int) {
	swap := IntSlice(data)
	cocktailSort(swap)
}

// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(n^2)
// 最优时间复杂度 ---- 如果序列在一开始已经大部分排序过的话,会接近O(n)
// 平均时间复杂度 ---- O(n^2)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 稳定
func cocktailSort(data Swap) {
	var left = 0
	var right = data.Len() - 1
	for left < right {
		for i := left; i < right; i++ {
			if data.Less(i+1, i) {
				data.Swap(i+1, i)
			}
		}
		right--
		for j := right; j > left; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
		left++
	}
}

// SelectionSortInts 选择排序
func SelectionSortInts(data []int) {
	swap := IntSlice(data)
	selectionSort(swap)
}

// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(n^2)
// 最优时间复杂度 ---- O(n^2)
// 平均时间复杂度 ---- O(n^2)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 不稳定
func selectionSort(data Swap) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		if min != i {
			data.Swap(min, i)
		}
	}
}

// InsertionSortInts 插入排序
// 分类 ------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- 最坏情况为输入序列是降序排列的,此时时间复杂度O(n^2)
// 最优时间复杂度 ---- 最好情况为输入序列是升序排列的,此时时间复杂度O(n)
// 平均时间复杂度 ---- O(n^2)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 稳定
func InsertionSortInts(data []int) {
	n := len(data)
	for i := 1; i < n; i++ {
		get := data[i]
		j := i - 1
		for j >= 0 && data[j] > get {
			data[j+1] = data[j]
		}
		data[j+1] = get
	}
}

// InsertionSortDichotomyInts 二分插入排序
// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(n^2)
// 最优时间复杂度 ---- O(nlogn)
// 平均时间复杂度 ---- O(n^2)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 稳定
func InsertionSortDichotomyInts(data []int) {
	n := len(data)
	for i := 1; i < n; i++ {
		get := data[i]
		left := 0
		right := i - 1
		for left <= right {
			mid := (left + right) / 2
			if data[mid] > get {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		for j := i - 1; j >= left; j-- {
			data[j+1] = data[j]
		}
		data[left] = get
	}
}

// ShellSortInts 希尔排序
// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- 根据步长序列的不同而不同。已知最好的为O(n(logn)^2)
// 最优时间复杂度 ---- O(n)
// 平均时间复杂度 ---- 根据步长序列的不同而不同。
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 不稳定
func ShellSortInts(data []int) {
	n := len(data)
	h := 0
	for h <= n {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			j := i - h
			get := data[i]
			for j >= 0 && data[j] > get {
				data[j+h] = data[j]
				j = j - h
			}
			data[j+h] = get
		}
		h = (h - 1) / 3
	}
}

// MergeSortInts 归并排序
// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(nlogn)
// 最优时间复杂度 ---- O(nlogn)
// 平均时间复杂度 ---- O(nlogn)
// 所需辅助空间 ------ O(n)
// 稳定性 ------------ 稳定
func MergeSortInts(data []int, left, mid, right int) { // 合并两个已排好序的数组A[left...mid]和A[mid+1...right]
	len := right - left + 1
	temp := make([]int, len)
	index := 0
	i := left
	j := mid + 1
	for i <= mid && j <= right {
		if data[i] <= data[j] {
			temp[index] = data[i]
			i++
		} else {
			temp[index] = data[j]
			j++
		}
		index++
	}

	for i <= mid {
		temp[index] = data[i]
		index++
		i++
	}

	for j <= right {
		temp[index] = data[j]
		index++
		j++
	}

	for k := 0; k < len; k++ {
		data[left] = temp[k]
		left++
	}
}

// MergeSortRecursionInts 递归实现的归并排序(自顶向下)
func MergeSortRecursionInts(deta []int, left, right int) {
	if left == right {
		return
	}
	mid := (left + right) / 2
	MergeSortRecursionInts(deta, left, mid)
	MergeSortRecursionInts(deta, mid, right)
	MergeSortInts(deta, left, mid, right)
}

// MergeSortIterationInts 非递归(迭代)实现的归并排序(自底向上)
func MergeSortIterationInts(data []int) {
	var left, mid, right int
	n := len(data)
	for i := 1; i < n; i *= 2 {
		left = 0
		for left+i < n {
			mid = left + i - 1
			if mid+i < n {
				right = mid + i
			} else {
				right = n - 1
			}
			MergeSortInts(data, left, mid, right)
			left = right + 1
		}
	}
}

// HeapSortInts 堆排序
func HeapSortInts(data []int) {
	swap := IntSlice(data)
	heapSort(swap)
}

// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(nlogn)
// 最优时间复杂度 ---- O(nlogn)
// 平均时间复杂度 ---- O(nlogn)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 不稳定
func heapSort(data Swap) {
	heapSize := buildHeap(data)
	for heapSize > 1 {
		heapSize--
		data.Swap(0, heapSize)
		heapify(data, 0, heapSize)
	}
}

func heapify(data Swap, i, size int) {
	leftChild := 2*i + 1
	rightChild := 2*i + 2
	max := i
	if leftChild < size && data.Less(max, leftChild) {
		max = leftChild
	}
	if rightChild < size && data.Less(max, rightChild) {
		max = rightChild
	}
	if max != i {
		data.Swap(i, max)
		heapify(data, max, size)
	}
}

func buildHeap(data Swap) int {
	heapSize := data.Len()
	for i := heapSize/2 - 1; i >= 0; i-- {
		heapify(data, i, heapSize)
	}
	return heapSize
}

// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(n^2)
// 最优时间复杂度 ---- O(nlogn)
// 平均时间复杂度 ---- O(nlogn)
// 所需辅助空间 ------ O(logn~n)
// 稳定性 ------------ 不稳定

// QuickSort 快排
func QuickSort(data Swap) {
	n := data.Len()
	for i := 0; i < n; i++ {
		quicksort(data, i, n-1)
	}
}

func quicksort(data Swap, left, right int) {
	if left > right {
		return
	}
	tempI := left
	tempJ := right
	for tempI != tempJ {
		for (!data.Less(tempJ, left)) && tempI < tempJ {
			tempJ--
		}
		for (!data.Less(left, tempI)) && tempI < tempJ {
			tempI++
		}
		if tempI < tempJ {
			data.Swap(tempI, tempJ)
		}
	}
	data.Swap(left, tempI)

	quicksort(data, left, tempI-1)
	quicksort(data, tempI+1, right)
}
