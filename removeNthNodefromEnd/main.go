// Given the head of a linked list, remove the nth node from the end of the list and return its head.

// Example 1:

// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Example 2:

// Input: head = [1], n = 1
// Output: []

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLL(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	head := &ListNode{Val: data[0]}
	temp := head
	for i := 1; i < len(data); i++ {
		temp.Next = &ListNode{Val: data[i]}
		temp = temp.Next
	}
	return head
}

func printList(ll *ListNode) {
	temp := ll
	for temp != nil {
		fmt.Println("value is:", temp.Val)
		temp = temp.Next
	}
}

// func removeNthNode(head *ListNode, n int) *ListNode {
// 	res := &ListNode{0, head}
// 	lead := res
// 	for i := 0; i <= n; i++ {
// 		lead = lead.Next
// 	}

// 	curr := res
// 	for lead != nil {
// 		curr = curr.Next
// 		lead = lead.Next
// 	}
// 	curr.Next = curr.Next.Next
// 	return res.Next

// }

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	head := createLL(data)
	printList(head)
	head1 := removeNthNode(head, 3)
	printList(head1)
}

func removeNthNode(head *ListNode, k int) *ListNode {
	res := &ListNode{0, head}
	lead := res
	for i := 0; i <= k; i++ {
		lead = lead.Next
	}

	curr := res
	for lead != nil {
		lead = lead.Next
		curr = curr.Next
	}
	fmt.Println("curr val is:", curr.Val)
	curr.Next = curr.Next.Next
	return res.Next
}
