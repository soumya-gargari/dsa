// Given the head of a linked list, rotate the list to the right by k places.
// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	head := createLL(data)
	printList(head)
	head1 := rotateRight(head, 2)
	fmt.Println("after rotating right")
	printList(head1)
}

// time complexity : O(n)
// space complexity : O(1)
// func rotateRight(head *ListNode, k int) *ListNode {
// 	if head == nil || head.Next == nil || k == 0 {
// 		return head
// 	}
// 	length := 1
// 	temp := head
// 	for temp.Next != nil {
// 		length += 1
// 		temp = temp.Next
// 	}
// 	temp.Next = head
// 	fmt.Println("length is:", length)
// 	k %= length
// 	res := head
// 	for i := 0; i < length-k-1; i++ {
// 		res = res.Next
// 	}
// 	fmt.Println("partition value is:", res.Val)
// 	newHead := res.Next
// 	res.Next = nil
// 	fmt.Println("After partition")
// 	printList(newHead)
// 	//temp.Next = head
// 	return newHead

// }

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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return nil
	}
	temp := head
	length := 1
	for temp.Next != nil {
		length += 1
		temp = temp.Next

	}
	temp.Next = head
	k %= length
	res := head
	for i := 0; i < length-k-1; i++ {
		res = res.Next
	}
	fmt.Println("partition val is:", res.Val)
	newHead := res.Next
	res.Next = nil
	return newHead

}
