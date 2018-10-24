package mutils

func Int10Ints(num int) []int {
	return IntTInts(num, 10)
}

func IntTInts(num int, base uint8) []int {
	var ret []int
	for num > 0 {
		ret = append(ret, num%int(base))
		num = num / int(base)
	}
	ReversInts(ret)
	return ret
}

func Ints10Int(mods []int) int {
	return IntsTInt(mods, 10)
}

func IntsTInt(mods []int, base uint8) int {
	if len(mods) == 0 {
		return 0
	}
	ret := 0
	for _, mod := range mods {
		ret = ret*int(base) + mod
	}
	return ret
}
