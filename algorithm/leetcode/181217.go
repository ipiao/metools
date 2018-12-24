package leetcode

// 无重复字节的最长字串
// 给定一个字符串。，请你找出其中不含有重复字符串的最长字串的长度

func lengthOfLongestSubstring(s string) int {
	var bMap = make(map[rune]int)
	var startInd = 0
	var maxLen = 0
	for i, r := range s {
		ind, ok := bMap[r]
		if ok { // 字符已经存在
			l := i - startInd
			if l > maxLen {
				maxLen = l
			}
			startInd = ind + 1
		}
		bMap[r] = i
	}
	if maxLen == 0 {
		maxLen = len([]rune(s))
	}
	return maxLen
}
