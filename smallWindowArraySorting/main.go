// To find the smallest window in an array that needs to be sorted in order to make the entire array sorted, you can follow these steps in Go:

// Iterate through the array from left to right and find the first element that is out of order (smaller than the previous element).
// Similarly, iterate through the array from right to left and find the first element that is out of order (larger than the previous element).
// Find the minimum and maximum elements within the range identified by the above steps.
// Expand the range identified above to include any elements between the minimum and maximum that are out of order.
// The range identified is the smallest window that needs to be sorted.
// example : [1, 3, 5, 2, 6, 4, 8]
// here in the above example: from 3 to 6, if all values we will sort. whole array will be sorted, that is the expectation

package main

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Find the left boundary of the unsorted subarray
	left := 0
	for left < n-1 && nums[left] <= nums[left+1] {
		left++
	}

	// The array is already sorted
	if left == n-1 {
		return 0
	}

	// Find the right boundary of the unsorted subarray
	right := n - 1
	for right > 0 && nums[right] >= nums[right-1] {
		right--
	}

	// Find the minimum and maximum elements within the unsorted subarray
	fmt.Println("left and right pointer value is:", left, right)

	// left = 2 (2nd index)
	// right = 5 (5th index)
	minUnsorted, maxUnsorted := nums[left], nums[right]
	// minunsorted = 5
	// maxunsorted = 4
	for i := left; i <= right; i++ {
		if nums[i] < minUnsorted {
			minUnsorted = nums[i]
		}
		if nums[i] > maxUnsorted {
			maxUnsorted = nums[i]
		}
	}
	fmt.Println("min unsorted and max unsorted is:", minUnsorted, maxUnsorted)
	// minunsorted = 2
	// maxunsorted = 6
	// Expand the subarray to include any out of order elements
	for left >= 0 && nums[left] > minUnsorted {
		left -= 1
	}
	for right < n && nums[right] < maxUnsorted {
		right += 1
	}

	fmt.Println(right, left)
	return right - left - 1
}

func main() {
	arr := []int{1, 3, 5, 2, 6, 4, 8}
	fmt.Println("Length of smallest window to be sorted:", findUnsortedSubarray(arr)) // Output: 4
}
