package leetcode

import (
	"log"
	"testing"
)

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}
	m1 := make(map[string]int)
	for _, word := range words {
		m1[word] += 1
	}
	ret := []int{}
	sl := len(words[0])

	for i := 0; i < sl; i++ {
		cm := make(map[string]int)
		c := 0
		ind := i
		for st := i; st < len(s)-len(words)*sl; {
			log.Println(i, st, ind)
			w := s[ind : ind+sl]
			log.Println(w)
			log.Println(cm)
			if wc, ok := m1[w]; ok {
				if cm[w] == wc {
					if cm[s[st:st+sl]] > 0 {
						cm[s[st:st+sl]] -= 1
						c--
					}
					st += sl
				} else {
					cm[w]++
					ind += sl
					if c == len(words)-1 {
						ret = append(ret, st)
						if cm[s[st:st+sl]] > 0 {
							cm[s[st:st+sl]] -= 1
						}
						st += sl
					} else {
						c++
					}
				}
			} else {
				c = 0
				ind += sl
				st = ind
				cm = make(map[string]int)
			}
		}
	}
	return ret
}

func TestFindSub(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	ss := []string{"word", "good", "best", "good"}
	t.Log(findSubstring(s, ss))
}
