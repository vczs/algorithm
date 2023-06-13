package two_sum

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func Test(t *testing.T) {
	t.Run("暴力枚举:", testTwoSumEnumeration)
	t.Run("哈希表:", testtwoSumHash)
}

// 测试暴力枚举方法
func testTwoSumEnumeration(t *testing.T) {
	nums := []int{6, 2, 5, 4, 3}
	res := twoSumEnumeration(nums, 8)
	t.Logf("两数下标分别是：%v", res)
}

// 测试哈希表方法
func testtwoSumHash(t *testing.T) {
	nums := []int{1, 2, 5, 4, 3}
	res := twoSumHash(nums, 9)
	t.Logf("两数下标分别是：%v", res)
}
