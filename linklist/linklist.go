package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linklist struct {
	head   *node
	length int
}

func (ll *linklist) addAtFirst(val int) {
	newNode := &node{data: val}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	newNode.next = ll.head
	ll.head = newNode
	ll.length++
}

func (ll *linklist) addAtLast(val int) {
	newNode := &node{data: val}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	temp := ll.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	ll.length++
}

func displayVal(ll linklist) {
	temp := ll.head
	for temp != nil {
		fmt.Println("value is:", temp.data)
		temp = temp.next
	}
}

func detectAndRemoveLoop(ll *linklist) {
	if ll.head == nil {
		fmt.Println("linklist is empty")
		return
	}
	var isLoop bool
	slow := ll.head
	fast := ll.head.next
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			fmt.Printf("loop detected, removing loop")
			fast.next.next = nil
			isLoop = true
			break
		}
	}
	if isLoop {
		fmt.Println("after removing loop")
		displayVal(*ll)
	}
}

func joinLinklists(l1 *linklist, l2 *linklist) *linklist {
	l3 := &linklist{}
	if l1.head == nil && l2.head == nil {
		return l3
	}
	if l1.head == nil {
		return l2
	} else if l2.head == nil {
		return l1
	}
	l3.head = l1.head
	temp := l3.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = l2.head
	return l3
}

func main() {
	ll := &linklist{head: &node{data: 40}}
	l2 := &linklist{}
	ll.addAtFirst(30)
	ll.addAtFirst(20)
	ll.addAtFirst(10)
	ll.addAtLast(50)
	displayVal(*ll)
	l2.addAtLast(100)
	l2.addAtLast(120)
	l2.addAtLast(130)
	l2.addAtLast(140)
	l3 := joinLinklists(ll, l2)
	fmt.Println("after joining 2 linklists")
	displayVal(*l3)
	n1 := &node{data: 10}
	n2 := &node{data: 20}
	n3 := &node{data: 40}
	n4 := &node{data: 80}
	n5 := &node{data: 90}
	l4 := &linklist{head: n1}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	n5.next = n3
	detectAndRemoveLoop(l4)
}
