package leetcode

import (
	"fmt"
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

func reverseString(s string) string {
	bs := []byte(s)
	reversBytes(bs)
	return string(bs)
}

func reversBytes(bs []byte) {
	if len(bs) <= 1 {
		return
	}
	head := bs[0]
	reversBytes(bs[1:])
	copy(bs[:len(bs)-1], bs[1:])
	bs[len(bs)-1] = head
}

func reversInts(bs []int) {
	if len(bs) <= 1 {
		return
	}
	head := bs[0]
	reversInts(bs[1:])
	copy(bs[:len(bs)-1], bs[1:])
	bs[len(bs)-1] = head
}

func ints10Int(num int) []int {
	var ret []int
	for num > 0 {
		ret = append(ret, num%10)
		num = num / 10
	}

	return ret
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
