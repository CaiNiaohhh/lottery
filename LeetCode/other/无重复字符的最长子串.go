package main

func lengthOfLongestSubstring(s string) int {
	// 特判
	length := len(s)
	if length < 2 {
		return length
	}
	// 定义两个指针和一个map
	st, en, maxLen := 0, 1, 1
	Map := make(map[byte]bool)
	Map[s[st]] = true
	for en < length {
		if _, ok := Map[s[en]]; ok {
			for st <= en && Map[s[en]] {
				delete(Map, s[st])
				st += 1
			}
			Map[s[en]] = true
		} else {
			Map[s[en]] = true
			maxLen = Max(maxLen, en - st + 1)
		}
		en += 1
	}
	return maxLen
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
