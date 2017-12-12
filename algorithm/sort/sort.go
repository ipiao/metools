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
	l := data.Len()
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
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
	l := data.Len()
	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		if min != i {
			data.Swap(min, i)
		}
	}
}
