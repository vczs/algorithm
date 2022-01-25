package longestString

func longestStringOne(s string) int {
	m := map[byte]int{} // 记录每个字符是否出现过
	n := len(s)
	t, r := -1, 0 // t从-1开始，相当于我们在字符串的左边界的左侧，还没有开始移动
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1]) // 删除已经判断的起始字符，将他的下一个字符变为起始字符
		}
		for t+1 < n && m[s[t+1]] == 0 {
			m[s[t+1]]++ // 不断地移动右指针
			t++
		}
		r = max(r, t+1-i) // 判断所有无重复子字符串 取其中最大值
	}
	return r
}

func longestStringTwo(s string) int {
	m := map[int]int{} // 记录字符上一次出现的位置
	for i := 0; i < 128; i++ {
		m[i] = -1 // 将所有字符出现位置初始值设为-1
	}
	r, start := 0, 0
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		index := s[i]
		start = max(start, m[int(index)]+1)
		r = max(r, i-start+1)
		m[int(index)] = i
	}
	return r
}

// 两数取最大值
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
