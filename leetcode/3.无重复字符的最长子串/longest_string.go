package longest_string

func lengthOfLongestSubstring(s string) int {
	length := len(s)

	m := map[byte]int{}
	maxLength := 0
	j := 0

	for i := 0; i < length; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for j < length && m[s[j]] == 0 {
			m[s[j]]++
			j++
		}
		maxLength = max(j-i, maxLength)
	}

	return maxLength
}

// 两数取最大值
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
