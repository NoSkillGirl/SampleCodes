// Given a sorted and rotated array, find if there is a pair with a given sum
// Input: arr[] = {11, 15, 6, 8, 9, 10}, x = 16
// Output: true
// There is a pair (6, 10) with sum 16
// Input: arr[] = {11, 15, 26, 38, 9, 10}, x = 45
// Output: false
// There is no pair with sum 45.

package main

import "fmt"

func main() {
	arr := [6]int{11, 15, 6, 8, 9, 10}
	sum := 16
	findPair(arr, sum)

	arr2 := [6]int{11, 15, 26, 38, 9, 10}
	findPair(arr2, 45)
}

func findPair(arr [6]int, sum int) {
	n := len(arr)
	var smallestIndex int
	var largestIndex int
	smallestIndex = 0
	largestIndex = n - 1
	// Find the pivot
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			smallestIndex = i + 1
			largestIndex = i
			break
		}
	}

	i := smallestIndex
	j := largestIndex

	//infinate loop
	for {
		//if sum is found
		if arr[i]+arr[j] == sum {
			fmt.Println("There is a pair: (", arr[i], ",", arr[j], ")")
			break
			//if sum is smaller then given sum
		} else if arr[i]+arr[j] < sum {
			if i < n-1 {
				i = i + 1
			} else {
				i = 0
			}
			//if sum is greater then given sum
		} else {
			if j > 0 {
				j = j - 1
			} else {
				j = n - 1
			}
		}
		//if the indexs meet
		if i == j {
			fmt.Println("No pair found")
			break
		}
	}
}
