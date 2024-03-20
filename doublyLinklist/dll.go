package main

import "fmt"

type node struct {
	prev *node
	next *node
	data int
}

type doublyLinklist struct {
	head *node
	tail *node
}

func (dll *doublyLinklist) AddAtFirst(val int) {
	newNode := &node{data: val}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}
	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode
}

func (dll *doublyLinklist) AddAtLast(val int) {
	newNode := &node{data: val}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}
	dll.tail.next = newNode
	newNode.prev = dll.tail
	dll.tail = newNode
}

func displayForwards(dll *doublyLinklist) {
	temp := dll.head
	for temp != nil {
		fmt.Println("value from starting position is:", temp.data)
		temp = temp.next
	}
}

func displayBackwards(dll *doublyLinklist) {
	temp := dll.tail
	for temp != nil {
		fmt.Println("value from end position is:", temp.data)
		temp = temp.prev
	}
}

func main() {
	dll := &doublyLinklist{}
	dll.AddAtFirst(10)
	dll.AddAtFirst(9)
	dll.AddAtFirst(3)
	dll.AddAtLast(11)
	displayForwards(dll)
	fmt.Println()
	displayBackwards(dll)
}
