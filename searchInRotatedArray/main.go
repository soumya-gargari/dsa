// There is an integer array nums sorted in ascending order (with distinct values).

// Prior to being passed to your function, nums is possibly rotated at an unknown
//pivot index k (1 <= k < nums.length) such that the resulting array is
//[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
// For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

// Given the array nums after the possible rotation and an integer target,
// return the index of target if it is in nums, or -1 if it is not in nums.

// You must write an algorithm with O(log n) runtime complexity.

package main

import "fmt"

func main() {
	data := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("found at index:", searchRotatedArray(data, 4))
}

// Time Complexity: The time complexity is O(log⁡n)
// since we're performing a binary search over the elements of the array.

// Space Complexity: The space complexity is O(1)
// func searchRotatedArray(data []int, target int) int {
// 	low := 0
// 	high := len(data) - 1
// 	for low <= high {
// 		mid := (low + high) / 2
// 		if data[mid] == target {
// 			return mid
// 		}
// 		if data[low] <= data[mid] {
// 			if data[low] <= target && target < data[mid] {
// 				high = mid - 1
// 			} else {
// 				low = mid + 1
// 			}
// 		} else {
// 			if data[mid] < target && target <= data[high] {
// 				low = mid + 1
// 			} else {
// 				high = mid - 1
// 			}
// 		}
// 	}
// 	return -1
// }

func searchRotatedArray(data []int, target int) int {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == target {
			return mid
		}
		if data[low] <= data[mid] {
			if data[low] <= target && target < data[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if data[mid] < target && target <= data[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
