package longest_palindromic_substring

func longestPalindrome(s string) string {
	length := len(s)

	if length < 2 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < length; {
		if length-i < end/2 {
			break
		}
		left, right := i, i
		for right < length-1 && s[right] == s[right+1] {
			right++
		}
		i = right + 1
		for left > 0 && right < length-1 && s[left-1] == s[right+1] {
			right++
			left--
		}
		if right-left+1 > end {
			start, end = left, right-left+1
		}
	}
	return s[start : start+end]
}
