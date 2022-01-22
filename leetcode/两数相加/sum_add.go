package sum_add

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请将两个数相加，并以相同形式返回一个 表示和 的链表
*/

/*
由于输入的两个链表都是逆序存储数字的位数的，因此两个链表中同一位置的数字可以直接相加。
我们同时遍历两个链表，逐位计算它们的和，并与当前位置的进位值相加。
如果两个链表的长度不同，则可以认为长度短的链表的后面有若干个 0 。
此外，如果链表遍历结束后，carry>0，还需要在答案链表的后面附加一个节点，值为carry。

时间复杂度：O(max(m,n))。其中 m 和 n 分别为两个链表的长度。我们要遍历两个链表的全部位置，而处理每个位置只需要 O(1) 的时间。
空间复杂度：O(1)。注意返回值不计入空间复杂度。
*/
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
		sum := n1 + n2 + carry      // 链表同位置节点值的和加上进位值
		sum, carry = sum%10, sum/10 // 计算同位置结果值 和 进位值
		if result != nil {
			temp.Next = &ListNode{Val: sum}
			temp = temp.Next // 辅助节点游标进一位
		} else {
			result = &ListNode{Val: sum}
			temp = result // 将只有头节点的链表赋值给辅助节点
		}
	}
	if carry > 0 { // 末位有进位 链表增加一个节点
		temp.Next = &ListNode{Val: carry}
	}
	return result
}
