package search_median

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func Test(t *testing.T) {
	t.Run("寻找两个正序数组的中位数：", testFindMedianSortedArrays)
}

func testFindMedianSortedArrays(t *testing.T) {
	res := findMedianSortedArrays([]int{}, []int{1})
	t.Logf("中位数是：%v", res)
}
