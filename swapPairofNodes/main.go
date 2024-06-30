package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	head := createLL(data)
	printList(head)
	head2 := swapPairs(head)
	fmt.Println("After swapping the values")
	printList(head2)
}

func swapPairs(head *ListNode) *ListNode {
	var prev, curr, newList *ListNode = nil, head, head.Next
	for curr != nil && curr.Next != nil {
		// values to be swapped
		firstNode, secondNode := curr, curr.Next

		// swapping the value
		firstNode.Next = secondNode.Next // 1 pointing next to 3
		secondNode.Next = firstNode      // means 2 will be pointing to 1

		if prev != nil {
			prev.Next = secondNode
		}

		// updating prev and current value
		prev = firstNode
		curr = firstNode.Next
	}
	return newList
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
