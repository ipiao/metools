package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	ret := twoSum2([]int{3, 2, 3}, 6)
	t.Log(ret)
}

func TestValid(t *testing.T) {
	t.Log(isValidSudoku([][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}

func TestReverseString(t *testing.T) {
	t.Log(reverseString("hello"))
}

func TestFirstUniqChar(t *testing.T) {
	i := firstUniqChar("leetcode")
	t.Log(i)
}

func TestIntToByte(t *testing.T) {
	for i := 0; i < 256; i++ {
		bs := IntToByte(i)
		t.Log(i, "==", string(bs))
	}

	i := ByteToInt('c')
	t.Log(i)
}

func TestMyAtoi(t *testing.T) {
	n := myAtoi("  a124551")
	t.Log(n)
}
