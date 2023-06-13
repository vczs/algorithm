package sum_add

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func Test(t *testing.T) {
	t.Run("两数相加:", testSumAdd)
}

// 测试两数相加
func testSumAdd(t *testing.T) {
	l1, l2 := createData()
	res := sumAdd(l1, l2)
	for {
		t.Log(res.Val)
		if res.Next == nil {
			break
		}
		res = res.Next
	}
}

// 构造用例
func createData() (*ListNode, *ListNode) {
	l1 := &ListNode{Val: 2}
	add(l1, 4)
	add(l1, 3)
	l2 := &ListNode{Val: 5}
	add(l2, 6)
	add(l2, 4)
	return l1, l2
}

// 链表添加元素
func add(ln *ListNode, num int) {
	temp := ln
	for {
		if temp.Next == nil {
			break
		}
		temp = ln.Next
	}
	temp.Next = &ListNode{Val: num}
}
