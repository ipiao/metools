package cal

import (
	"log"
	"testing"
)

func TestCal(t *testing.T) {
	bn := []int{2, 3, 4, 5}
	cn := bn
	dn := make([]int, len(bn))
	dn = bn
	for i := range cn {
		cn[i]++
		dn[i]++
	}
	t.Log(bn)
}

func TestAutoAdd(t *testing.T) {
	ln := []int{2, 3, 4}
	bn1 := []int{0, 0, 0}
	flag := true
	for flag {
		flag = autoAdd(bn1, ln)
		t.Log(flag, bn1)
	}
}

func TestUgly(t *testing.T) {
	ug := NewUgly([]int{2, 3, 5})
	n := ug.Get(1500)
	t.Log(n)
	t.Log(sumLen(ug.ranks))
	t.Log(ug.ranks[len(ug.ranks)-1])
}

func TestInt(t *testing.T) {
	log.Println(int32(859963392))
}
