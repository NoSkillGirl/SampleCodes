// The selection sort algorithm sorts an array by repeatedly finding the 
// minimum element (considering ascending order) from unsorted part and 
// putting it at the beginning. The algorithm maintains two subarrays in a given array.

// 1) The subarray which is already sorted.
// 2) Remaining subarray which is unsorted.

// In every iteration of selection sort, the minimum element (considering ascending 
// order) from the unsorted subarray is picked and moved to the sorted subarray.

// Time Complexity: O(n2) as there are two nested loops.
// Auxiliary Space: O(1)

package main

import (
	"fmt"
)

func main() {
	arr := [5]int{64, 25, 12, 22, 11}
	selectionSort(arr)
}

func selectionSort (arr[5] int){
	fmt.Println("array: ", arr)
	n := len(arr)
	var minIdx int
	for i:=0; i<n-1; i++ {  
        minIdx = i  
        for j:=i+1; j<n; j++ {  
        	if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}  
	fmt.Println("sorted array: ", arr)
}
