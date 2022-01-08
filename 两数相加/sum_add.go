package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 2}
	add(l1, 4)
	add(l1, 3)
	l2 := &ListNode{Val: 5}
	add(l2, 6)
	add(l2, 4)
	result := addTwoNumbers(l1, l2)
	for {
		fmt.Print(result.Val)
		if result.Next == nil {
			break
		}
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var temp *ListNode
	var result *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if result != nil {
			temp.Next = &ListNode{Val: sum}
			temp = temp.Next
		} else {
			result = &ListNode{Val: sum}
			temp = result
		}
	}
	if carry > 0 {
		temp.Next = &ListNode{Val: carry}
	}
	return result
}
