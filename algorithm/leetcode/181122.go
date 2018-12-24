package leetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	v1 := sumstring(s1)
	i := len(s1)
	v2 := sumstring(s2[:i])
	for i <= len(s2) {
		if i > len(s1) {
			v2 += int(s2[i-1]) - int(s2[i-len(s1)-1])
		}
		if v1 == v2 {
			if p(s1, s2[i-len(s1):i]) {
				return true
			}
		}
		i++
	}
	return false
}

func sumstring(s string) int {
	sum := 0
	for i := range s {
		sum += int(s[i])
	}
	return sum
}

func p(b1, b2 string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := range b1 {
		m1[b1[i]] += 1
		m2[b2[i]] += 1
	}
	for b, i := range m1 {
		if m2[b] != i {
			return false
		}
	}
	return true
}
