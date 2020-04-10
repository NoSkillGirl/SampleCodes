//Bubble Sort is the simplest sorting algorithm that works by repeatedly
// swapping the adjacent elements if they are in wrong order.

// Worst and Average Case Time Complexity: O(n*n). Worst case occurs when array is reverse sorted.
// Best Case Time Complexity: O(n). Best case occurs when array is already sorted.
// Auxiliary Space: O(1)

package main

import (
	"fmt"
)

func main() {
	arr := [5]int{64, 25, 12, 22, 11}
	bubbleSort(arr)
}

func bubbleSort(arr [5]int) {
	fmt.Println("array: ", arr)
	n := len(arr)
	var swapped bool
	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
	fmt.Println("sorted array: ", arr)
}
