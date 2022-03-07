package main

func lengthOfLongestSubstring(s string) int {
	// 特判
	length := len(s)
	if length < 2 {
		return length
	}
	// 定义两个指针和一个map
	st, en, maxLen := 0, 1, 1
	Map := make(map[byte]int)
	// Map := map[byte]bool{}
	Map[s[st]] = 1
	for en < length {
		if _, ok := Map[s[en]]; ok && Map[s[en]] > 0 {
			for st <= en && Map[s[en]] > 0 {
				Map[s[st]] -= 1
				st += 1
			}
			Map[s[en]] += 1
		} else {
			Map[s[en]] += 1
			maxLen = Max(maxLen, (en - st + 1))
		}
		en += 1
	}
	return maxLen
}

func Max(x int, y int) (z int) {
	if x > y {
		z = x
	} else {
		z = y
	}
	return z
}
