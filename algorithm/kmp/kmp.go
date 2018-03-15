package kmp

func StringIndex(s, p string) int {
	return kmpIndex([]byte(s), []byte(p))
}

func kmpIndex(s, p []byte) int {
	i := 0
	j := 0
	next := getNext(p)
	for i < len(s) && j < len(p) {
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(p) {
		return i - j
	}
	return -1
}

func getNext(ms []byte) []int {
	length := len(ms)
	next := make([]int, length)
	next[0] = -1
	k := -1
	j := 0
	for j < length-1 {
		if k == -1 || ms[j] == ms[k] {
			j++
			k++
			if ms[j] != ms[k] {
				next[j] = k
			} else {
				next[j] = next[k]
			}
		} else {
			k = next[k]
		}
	}
	return next
}
