package main

import "fmt"

func partition(arr []int, low, high int) int {
	i := low
	j := high
	pivot := arr[i]
	// if len(arr) < 2 {
	// 	return arr[0]
	// }
	// keeping i< j as we should have atleast 2 elements in array,
	for i < j {
		// if arr[i] is less thgan pivot value, increment i as pivot will be middle value of array
		// in best case, so left side values of pivot will be lesser, right side values of pivot will be greater
		for arr[i] <= pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		// after incrementing and decrementing i and j, if i <j finally, then swap values
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// if above condition ends, that means time to return the pivot index as j and i overlapped
	// means j is smaller than i now
	// so time to change pivot value with j as j is smaller than i now
	arr[low], arr[j] = arr[j], arr[low]
	return j
}

func quicksort(a []int, low, high int) {
	if low < high {
		pivot := partition(a, low, high)
		fmt.Println("pivot val is:", pivot)
		quicksort(a, low, pivot-1)
		quicksort(a, pivot+1, high)
	}
}

func main() {
	data := []int{3, 2, 67, 45, 89, 1}
	high := len(data) - 1
	quicksort(data, 0, high)
	fmt.Println(data)
}
