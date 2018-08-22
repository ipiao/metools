package leetcode

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
)

// 1.给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
// 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
// O(n^2)
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//
func twoSum2(nums []int, target int) []int {
	onum := make([]int, len(nums))
	copy(onum, nums)

	// Onlgn
	sort.Ints(nums)

	var find = func(anum []int, dest int) int {
		start := 0
		end := len(anum) - 1
		var mid int
		for start < end-1 {
			mid = (start + end) / 2
			if anum[mid] == dest {
				return mid
			}
			if anum[mid] > dest {
				end = mid
			} else {
				start = mid
			}
		}
		if anum[start] == dest {
			return start
		}
		if anum[end] == dest {
			return end
		}
		return -1
	}

	for i := range nums {
		if cj := find(nums[i+1:], target-nums[i]); cj != -1 {
			fmt.Println("====", nums[i], cj)
			var ret = make([]int, 0)
			for j := range onum {
				if onum[j] == nums[i] || onum[j] == nums[i+1+cj] {
					ret = append(ret, j)
				}
			}
			return ret
		}
	}
	return nil
}

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]byte{}
	squars := [9][9]byte{}
	list := [9][9]byte{}
	for i, sb := range board {
		for j := range sb {
			rows[j][i] = sb[j]
			squars[(i/3)*3+(j/3)][(i%3)*3+(j%3)] = sb[j]
			list[i][j] = sb[j]
		}
	}
	for _, row := range list {
		if hasRepeted(row) {
			return false
		}
	}
	for _, row := range rows {
		if hasRepeted(row) {
			return false
		}
	}
	for _, row := range squars {
		if hasRepeted(row) {
			return false
		}
	}
	return true

}

func hasRepeted(bs [9]byte) bool {
	for i := 0; i < len(bs)-1; i++ {
		for j := i + 1; j < len(bs); j++ {
			if bs[i] != '.' && bs[i] == bs[j] {
				return true
			}
		}
	}
	return false
}

// 顺时针旋转 (x,y)->(y,-x)
// 如果原点为（n-1，0）
// 第一象限全部旋转到第四象限
// 然后在平移回到第一象限
// (x,y)–n>(y,n-x)
func rotate(matrix [][]int) {
	l := len(matrix)
	// 首先获得一块初始旋转区域和原点
	// 虚拟坐标系
	maxX := l / 2
	maxY := maxX
	if l%2 == 1 { //
		maxY += 1
	}

	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			// 4次旋转
			temp := matrix[i][j]
			matrix[i][j] = matrix[l-1-j][i]
			matrix[l-1-j][i] = matrix[l-1-i][l-1-j]
			matrix[l-1-i][l-1-j] = matrix[j][l-1-i]
			matrix[j][l-1-i] = temp
		}
	}
}

func firstUniqChar(s string) int {
	var baseArr [256]int
	for i := range s {
		tmp := int(s[i])
		baseArr[tmp] += 1
	}

	for k := range s {
		if baseArr[s[k]] == 1 {
			return k
		}
	}
	return -1
}

//整形转换成字节
func IntToByte(n int) byte {
	// tmp := int32(n)
	// bytesBuffer := bytes.NewBuffer([]byte{})
	// binary.Write(bytesBuffer, binary.BigEndian, tmp)
	// return bytesBuffer.Bytes()
	return byte(n)
}
func ByteToInt(b byte) int {
	// bs := make([]byte, 4)
	// bs[3] = b
	return int(b)
	// bytesBuffer := bytes.NewBuffer(bs)
	// var tmp int32
	// binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	// return int(tmp)
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if !(('0' <= s[i] && s[i] <= '9') || ('a' <= s[i] && s[i] <= 'z')) {
			i++
			continue
		}
		if !(('0' <= s[j] && s[j] <= '9') || ('a' <= s[j] && s[j] <= 'z')) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func myAtoi(str string) int {
	var max = math.MaxInt32
	bas := 1 // 乘子
	nbs := make([]int, 0)
	start := false

	for i := range str {
		b := str[i]
		if !start {
			if b == '-' {
				bas = -1
				start = true
			} else if b == '+' || ('0' <= b && b <= '9') {
				start = true
			}
		}

		if '0' <= b && b <= '9' {
			nbs = append(nbs, int(b-'0'))
		} else {
			break
		}
	}

	if bas == -1 {
		max += 1
	}
	ret := 0
	for i := 0; i < len(nbs); i++ {
		if ret > max/10 {
			ret = max
			break
		}
		ret = ret*10 + nbs[i]
	}
	return bas * ret
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	base := countAndSay(n - 1)
	i := 0
	count := 1
	ret := ""
	for j := 1; j < len(base); j++ {
		if base[j] == base[i] {
			count++
		} else {
			ret += fmt.Sprint(count, string(base[i]))
			i = j
			count = 1
		}
	}
	ret += fmt.Sprint(count, string(base[i]))
	return ret
}

func longestCommonPrefix(strs []string) string {
	base := strs[0]
	for i := 0; i < len(base); i++ {
		b := base[i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != b {
				return base[:i]
			}
		}
	}
	return base
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var ret *ListNode
	if l1.Val <= l2.Val {
		ret = l1
		l1 = l1.Next
		ret.Next = nil
	} else {
		ret = l2
		l2 = l2.Next
		ret.Next = nil
	}
	retLast := ret
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			retLast.Next = l1
			l1 = l1.Next
			retLast = retLast.Next
			retLast.Next = nil
		} else {
			retLast.Next = l2
			l2 = l2.Next
			retLast = retLast.Next
			retLast.Next = nil
		}
	}
	if l1 != nil {
		retLast.Next = l1
	}
	if l2 != nil {
		retLast.Next = l2
	}
	return ret
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preNode := head
	node := head.Next
	preNode.Next = nil
	nextNode := node.Next

	for nextNode != nil {
		node.Next = preNode
		preNode = node
		node = nextNode
		nextNode = node.Next
	}
	node.Next = preNode
	head = node
	return head
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList2(head.Next)
	if head.Next != nil {
		head.Next.Next = head
		head.Next = nil
	}
	return node
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	itNode := head
	for i := 1; i < n; i++ {
		itNode = itNode.Next
	}

	var npNode *ListNode
	nNode := head
	for itNode.Next != nil {
		itNode = itNode.Next
		npNode = nNode
		nNode = nNode.Next
	}
	if npNode == nil {
		return nNode.Next
	}
	npNode.Next = nNode.Next
	return head
}

func (node *ListNode) String() string {
	bs, _ := json.Marshal(node)
	return string(bs)
}

func isPalindromeList(head *ListNode) bool {
	node := head
	count := 0
	for node != nil {
		node = node.Next
		count++
	}
	if count <= 1 {
		return true
	}
	midNode := head
	for i := 0; i < (count)/2; i++ {
		midNode = midNode.Next
	}
	if count%2 == 1 {
		midNode = midNode.Next
	}
	midNode = reverseList(midNode)
	log.Println(midNode)
	log.Println(head)
	for midNode != nil {
		if head.Val != midNode.Val {
			return false
		}
		midNode = midNode.Next
		head = head.Next
	}
	return true
}

func hasCycle(node *ListNode) bool {
	if node == nil {
		return false
	}
	for node.Next != nil {
		if node.Next == node {
			return true
		}
		node = node.Next
	}
	return false
}
