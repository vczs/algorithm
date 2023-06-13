package longest_palindromic_substring

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}
func Test(t *testing.T) {
	t.Run("测试最长回文子串: ", testLongestPalindrome)
}
func testLongestPalindrome(t *testing.T) {
	res := longestPalindrome("babad")
	t.Logf("子串是: %v", res)
}
