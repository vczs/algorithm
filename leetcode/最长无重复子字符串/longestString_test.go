package longestString

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func Test(t *testing.T) {
	t.Run("无重复最长子字符串方法1：", testLongestStringOne)
	t.Run("无重复最长子字符串方法2：", testLongestStringTwo)
}

// 测试testLongestStringOne方法
func testLongestStringOne(t *testing.T) {
	r := longestStringOne("abcabcbb")
	t.Logf("方法1“abcabcbb”中无重复最长子字符串的长度是：%v", r)

}

// 测试testLongestStringTwo方法
func testLongestStringTwo(t *testing.T) {
	r := longestStringTwo("abcabcbb")
	t.Logf("方法2“abcabcbb”中无重复最长子字符串的长度是：%v", r)
}
