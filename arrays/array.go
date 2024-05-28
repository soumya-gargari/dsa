package main

import "fmt"

type array struct {
	data []int
}

func (a *array) Sum() {
	sum := 0
	for _, val := range a.data {
		sum += val
	}
	fmt.Println("sum of all elements are:", sum)
}

func (a *array) findLargestNum() {
	maxVal := a.data[0]
	for _, val := range a.data {
		if val > maxVal {
			maxVal = val
		}
	}
	fmt.Println("largest number within the array is:", maxVal)
}


// [3, 2, 4, 1]
func (a *array) find2ndLargestNum() {
	large1 := a.data[0]
	large2 := 0
	for i := 1; i < len(a.data); i++ {
		if large1 < a.data[i] {
			large2 = large1
			large1 = a.data[i]
		} else if large2 < a.data[i] {
			large2 = a.data[i]
		}
	}
	fmt.Println("2nd largest number within the array is:", large2)
}

func (a *array) binarySearch(value int) {
	first := 0
	last := len(a.data)
	middle := (first + last) / 2
	for first <= last {
		if value == a.data[middle] {
			fmt.Println("found the data at index:", middle)
			return
		}
		if value < a.data[middle] {
			last = last - 1
		} else {
			first = first + 1
		}
		middle = (first + last) / 2
	}
}

func main() {
	var a array
	data := []int{2, 6, 23, 43, 54}
	a.data = data
	a.Sum()
	a.findLargestNum()
	a.find2ndLargestNum()
	a.binarySearch(54)
}
