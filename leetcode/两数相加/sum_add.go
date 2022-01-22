package sum_add

type ListNode struct {
	Val  int
	Next *ListNode
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
