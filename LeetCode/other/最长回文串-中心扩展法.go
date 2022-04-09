package main
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	length := len(s)
	res := ""
	for i := 0; i < length; i ++ {
		res = maxLen(maxLen(calcS(s, i), calcD(s, i)), res)
	}
	return res

}
// 计算当前的回文串是中间单个字符
func calcS(s string, index int) string {
	left, right, length := index - 1, index + 1, len(s)
	for {
		if left < 0 || right >= length || s[left] != s[right] {
			return s[left + 1:right]
		}
		left, right = left - 1, right + 1
	}
}

// 计算当前的回文串是中间双字符
func calcD(s string, index int) string {
	left, right, length := index, index + 1, len(s)
	for {
		if left < 0 || right >= length || s[left] != s[right] {
			if left > right - 2 {
				return s[left:left]
			}
			return s[left + 1:right]
		}
		left, right = left - 1, right + 1
	}
}

func maxLen(s1 string, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	} else {
		return s2
	}
}
