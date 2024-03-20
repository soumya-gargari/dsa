package main

import "fmt"

type stack struct {
	array1 []int
}

type queue struct {
	array2 []int
}

// in stack, value gets appended at last while pushing
func (s *stack) push(val int) {
	s.array1 = append(s.array1, val)
}

// in stack, value gets poped up which is appended at last
func (s *stack) pop() int {
	poppedVal := s.array1[len(s.array1)-1]
	s.array1 = s.array1[:len(s.array1)-1]
	return poppedVal
}

// in queue, value gets appended at last while enqueuing
func (q *queue) enque(val int) {
	q.array2 = append(q.array2, val)
}

// in queue, value gets dequeued which is appended at 1st
func (q *queue) deque() int {
	dequedVal := q.array2[0]
	q.array2 = q.array2[1:]
	return dequedVal
}

func main() {
	s := &stack{array1: []int{2, 3, 4, 5}}
	s.push(6)
	s.push(9)
	s.pop()
	fmt.Println("stack values are:", s.array1)
	q := &queue{array2: []int{2, 3, 4, 5, 90, 89}}
	q.enque(100)
	q.enque(190)
	//q.deque()
	fmt.Println("queue values are:", q.array2)
	fmt.Println(q.array2[3:6])
}
