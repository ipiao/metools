package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return cbs2(candidates, 0, target)
}

func cbs2(nums []int, start, target int) [][]int {
	if target <= 0 || nums[start] > target {
		return nil
	}

	ret := [][]int{}

	if nums[start] == target {
		ret = append(ret, []int{nums[start]})
		return ret
	}

	for i := start + 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		r1s := cbs(nums, i, target-nums[start])
		for _, rs := range r1s {
			r := make([]int, len(rs)+1)
			r[0] = nums[start]
			copy(r[1:], rs)
			ret = append(ret, r)
		}
	}

	ret = append(ret, cbs2(nums, start+1, target)...)
	return ret
}
