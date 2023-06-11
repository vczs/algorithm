package longest_string

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func Test(t *testing.T) {
	t.Run("无重复最长子字符串：", testLongestString)
}

// 测试testLongestStringOne方法
func testLongestString(t *testing.T) {
	r := longestString("dvdf")
	t.Logf("“abcabcbb”中无重复最长子字符串的长度是：%v", r)
}
