package main

import (
	"fmt"
)

const ArraySize = 7

// hashtables are combo of array and linklist, time complexity is O(1)
// hashtable structure
type hashtable struct {
	array [ArraySize]*bucket
}

// bucket structure
type bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Init func initializes the hashtable structure
func Init() *hashtable {
	ht := &hashtable{}
	for i := 0; i < ArraySize; i++ {
		ht.array[i] = &bucket{}
	}
	return ht
}

// Insert
func (h *hashtable) Insert(val string) {
	index := hash(val)
	h.array[index].insert(val)
}

// Delete
func (h *hashtable) Delete(val string) {
	index := hash(val)
	h.array[index].delete(val)
}

// Search
func (h *hashtable) Search(val string) bool {
	index := hash(val)
	return h.array[index].search(val)
}

// insert
func (b *bucket) insert(val string) {
	newNode := &bucketNode{key: val}
	if b.head == nil {
		b.head = newNode
		return
	}
	newNode.next = b.head
	b.head = newNode

}

// delete
func (b *bucket) delete(val string) {
	if b.head.key == val {
		b.head = b.head.next
		return
	}
	temp := b.head
	for temp.next != nil {
		if temp.next.key == val {
			temp = temp.next.next
			return
		}
		temp = temp.next
	}
}

// search
func (b *bucket) search(val string) bool {
	temp := b.head
	for temp.next != nil {
		if temp.key == val {
			return true
		}
		temp = temp.next
	}
	return false
}

// hash func hashifies the string and return value
func hash(data string) int {
	sum := 0
	for _, val := range data {
		sum += int(val)
	}
	return (sum % ArraySize)
}
func main() {
	hashTable := Init()
	hashTable.Insert("soumya")
	hashTable.Insert("sanket")
	hashTable.Insert("chona")
	hashTable.Insert("joy")
	hashTable.Insert("akash")
	hashTable.Insert("sonu")
	fmt.Println(hashTable.Search("sonu"))
	hashTable.Delete("joy")
	fmt.Println(hashTable.Search("joy"))
}
