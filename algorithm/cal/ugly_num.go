package cal

import (
	"math"
	"sort"
)

// 要求输入一个n数输出第n个丑数。丑数是素因子只有2.3.5.7...。非常急，谢谢。

// Ugly 丑数，基础
// base 素数数组 可以是 2,3,5|3,5,7
type Ugly struct {
	base  []int   // 基础素因子
	bmul  int     // 基础素因子 乘积
	cbase [][]int // 素因子对应的第i次计算次数
	ranks [][]int // 素因子各次值排序
}

// NewUgly 创建一个丑数计算基础
func NewUgly(base []int) *Ugly {
	sort.Ints(base)
	u := &Ugly{
		base: base,
		bmul: mul(base),
	}
	u.calNext()
	return u
}

// 第n次计算
func (u *Ugly) calNext() {
	var n = len(u.cbase) + 1
	var bb = make([]int, len(u.base))
	for i, b := range u.base {
		bb[i] = maxc(pow(u.bmul, n), b)
	}
	u.cbase = append(u.cbase, bb)
	//
	bn := make([]int, len(bb))
	var rank []int
	min := pow(u.bmul, n-1)
	for autoAdd(bn, bb) {
		pm := powmul(u.base, bn)
		if pm > min && pm <= pow(u.bmul, n) {
			rank = append(rank, pm)
		}
	}
	sort.Ints(rank)
	u.ranks = append(u.ranks, rank)
}

// 多元排序自增
// 返回是否还有下一个
// ln下限
func autoAdd(bn, ln []int) bool {
	if bn[0] < ln[0] {
		bn[0]++
		return true
	}
	bn[0] = 0
	if len(bn) == 1 {
		return false
	}
	return autoAdd(bn[1:], ln[1:])
}

// 获取a以b为底的整数部分
func maxc(a, b int) int {
	return int(math.Log(float64(a)) / math.Log(float64(b)))
}

// 获取乘积
func mul(nums []int) int {
	ret := 1
	for _, n := range nums {
		ret *= n
	}
	return ret
}

// 获取数组所有长度
func sumLen(a [][]int) int {
	l := 0
	for _, aa := range a {
		l += len(aa)
	}
	return l
}

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	ret := 1
	for i := 0; i < n; i++ {
		ret *= x
	}
	return ret
}

func powmul(a, b []int) int {
	var c = make([]int, len(a))
	for i := range a {
		c[i] = pow(a[i], b[i])
	}
	return mul(c)
}

// Get 获取第n个丑数
func (u *Ugly) Get(n int) int {
	for sumLen(u.ranks) < n {
		u.calNext()
	}
	sl := 0
	ret := 0
	for i := range u.ranks {
		li := len(u.ranks[i])
		if sl+li >= n {
			ret = u.ranks[i][n-sl-1]
			break
		}
		sl += li
	}
	return ret
}
