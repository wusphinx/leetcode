package medium

/*
2. Add Two Numbers
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b := l1, l2
	c := ListNode{}
	cIter := &c
	ext := 0
	for a != nil && b != nil {
		val := a.Val + b.Val + ext
		if val >= 10 {
			ext = 1
			val = val - 10
		} else {
			ext = 0
		}
		tmp := &ListNode{Val: val}
		cIter.Next = tmp
		a = a.Next
		b = b.Next
		cIter = tmp
	}

	for a != nil {
		val := a.Val + ext
		if val >= 10 {
			ext = 1
			val = val - 10
		} else {
			ext = 0
		}
		tmp := &ListNode{Val: val}
		cIter.Next = tmp
		cIter = tmp
		a = a.Next

	}

	for b != nil {
		val := b.Val + ext
		if val >= 10 {
			ext = 1
			val = val - 10
		} else {
			ext = 0
		}
		tmp := &ListNode{Val: val}
		cIter.Next = tmp
		cIter = tmp
		b = b.Next
	}

	if ext > 0 {
		tmp := &ListNode{Val: ext}
		cIter.Next = tmp
	}

	return c.Next
}

func initListNode(src ...int) *ListNode {
	a := ListNode{}
	aa := &a
	for _, i := range src {
		tmp := &ListNode{Val: i}
		aa.Next = tmp
		aa = tmp
	}
	aa.Next = nil
	return a.Next
}

func getListNode(m *ListNode) (r []int) {
	s := m
	r = make([]int, 0)
	for s != nil {
		r = append(r, s.Val)
		s = s.Next
	}
	return
}
