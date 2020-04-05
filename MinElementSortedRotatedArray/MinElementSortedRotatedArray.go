// Find the minimum element in a sorted and rotated array
// A sorted array is rotated at some unknown point, find the minimum element in it.

// Following solution assumes that all elements are distinct.

// Input: {5, 6, 1, 2, 3, 4}
// Output: 1

package main

import "fmt"

func main() {
	arr := [6]int{5, 6, 1, 2, 3, 4}
	n := len(arr)
	minEle := findMin(arr, 0, n-1)
	fmt.Println("Min element: ", minEle)
}

func findMin(arr [6]int, low int, high int) int {
	// This condition is needed to
	// handle the case when array is not
	// rotated at all
	if high < low {
		return arr[0]
	}

	// If there is only one element left
	if high == low {
		return arr[low]
	}

	// Find mid
	mid := low + (high-low)/2

	// Check if element (mid+1) is minimum element. Consider
	// the cases like {3, 4, 5, 1, 2}
	if mid < high && arr[mid+1] < arr[mid] {
		return arr[mid+1]
	}

	// Check if mid itself is minimum element
	if mid > low && arr[mid] < arr[mid-1] {
		return arr[mid]
	}

	// Decide whether we need to go to left half or right half
	if arr[high] > arr[mid] {
		return findMin(arr, low, mid-1)
	}
	return findMin(arr, mid+1, high)
}
