package cal

import "github.com/ipiao/metools/algorithm/sort"

// 给定一个非负整数的列表，重新排列他们的顺序把他们组成一个最大的整数

// MaxNumInts 实现排序Swap接口
type MaxNumInts []int

func (s MaxNumInts) Less(i, j int) bool {
	ai := array(s[i])
	aj := array(s[j])
	return less(aj, ai)
}

func (s MaxNumInts) Swap(i, j int) { temp := s[i]; s[i] = s[j]; s[j] = temp }
func (s MaxNumInts) Len() int      { return len(s) }

func array(i int) []int {
	if i < 10 {
		return []int{i}
	}
	e := i / 10
	f := i % 10
	return append(array(e), f)
}

func less(ai, aj []int) bool {
	for k := 0; ; k++ {
		if k < len(ai) && k < len(aj) {
			if ai[k] != aj[k] {
				return ai[k] < aj[k]
			}
		} else {
			break
		}
	}
	if len(aj) > len(ai) {
		return less(ai, aj[len(ai):])
	} else if len(aj) < len(ai) {
		return less(ai[len(aj):], aj)
	}
	return false
}

func MaxNum(s []int) []int {
	ns := MaxNumInts(s)
	sort.QuickSort(ns)
	return s
}
