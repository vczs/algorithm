package longest_string

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func Test(t *testing.T) {
	t.Run("无重复最长子字符串：", testLengthOfLongestSubstring)
}

func testLengthOfLongestSubstring(t *testing.T) {
	res := lengthOfLongestSubstring("dvdf")
	t.Logf("“abcabcbb”中无重复最长子字符串的长度是：%v", res)
}
