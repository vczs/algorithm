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
	result := twoSumEnumeration(nums, 8)
	if result != nil {
		t.Logf("两数下标分别是：%v", result)
	} else {
		t.Logf("数组中没有符合条件的数")
	}
}

// 测试哈希表方法
func testtwoSumHash(t *testing.T) {
	nums := []int{1, 2, 5, 4, 3}
	result := twoSumHash(nums, 9)
	if result != nil {
		t.Logf("两数下标分别是：%v", result)
	} else {
		t.Logf("数组中没有符合条件的数")
	}
}
