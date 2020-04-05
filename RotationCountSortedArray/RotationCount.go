// Find the Rotation Count in Rotated Sorted array
// Consider an array of distinct numbers sorted in increasing order. 
// The array has been rotated (clockwise) k number of times. Given such an array, find the value of k.

// Input : arr[] = {15, 18, 2, 3, 6, 12}
// Output: 2
// Explanation : Initial array must be {2, 3,
// 6, 12, 15, 18}. We get the given array after 
// rotating the initial array twice.

// Time Complexity : O(n)
// Auxiliary Space : O(1)

package main

import "fmt"

func main(){
	arr := [6]int{12, 15, 18, 2, 3, 6}
	n := len(arr)
	count := countRotations(arr, 0, n-1)
	fmt.Println("Rotation count: ", count)
}

func countRotations (arr [6]int, low int, high int) int{
	// fmt.Println("low: ",low, " high : ", high)
	// This condition is needed to handle the case 
    // when the array is not rotated at all 
    if (high < low) {
		//fmt.Println("Rotation counnt : 0")
		return 0
	}

    // If there is only one element left 
    if high == low {
		// fmt.Println("Rotation counnt :", low)
		return low
	} 

	// Find mid 
	mid := low + (high - low)/2
	
	// Check if element (mid+1) is minimum element. 
    // Consider the cases like {3, 4, 5, 1, 2} 
    if mid < high && arr[mid+1] < arr[mid] {
		// fmt.Println("Rotation counnt :", mid+1)
		return mid+1
	}
	
	// Check if mid itself is minimum element 
    if mid > low && arr[mid] < arr[mid - 1]{
		// fmt.Println("Rotation counnt :", mid)
		return mid
	} 
       
    // Decide whether we need to go to left half or 
    // right half 
    if arr[high] > arr[mid]{
		return countRotations(arr, low, mid-1)
	}  
  
    return countRotations(arr, mid+1, high)
       
}