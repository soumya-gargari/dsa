package main

import (
	"container/heap"
	"fmt"
)

func main() {
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
}

type KthLargest struct {
	Heap *MinHeap
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	obj := KthLargest{
		Heap: h,
		K:    k,
	}
	for _, val := range nums {
		obj.Add(val)
	}
	return obj
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.Heap, val)
	if this.Heap.Len() > this.K {
		heap.Pop(this.Heap)
	}
	return (*this.Heap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
type MinHeap []int

func (h *MinHeap) Push(val any) {
	*h = append(*h, val.(int))
}

// Len() method should be there if you are using heap package for heap init method
func (h MinHeap) Len() int {
	return len(h)
}

// Less() method should be there if you are using heap package for heap init method
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap() method should be there if you are using heap package for heap init method
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Pop() any {
	old := *h
	length := len(*h)
	x := old[length-1]
	*h = old[0 : length-1]
	return x
}
