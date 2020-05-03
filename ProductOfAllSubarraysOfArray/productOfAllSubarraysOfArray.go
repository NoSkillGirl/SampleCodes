// Given an array of integers arr of size N, the task is to print products of all
// subarrays of the array.

// Examples:
// nput : arr[] = {10, 3, 7}
// Output : 30870
// Here, subarrays are [10], [10, 3], [10, 3, 7], [3], [3, 7], [7]
// Prodcuts are 10, 30, 210, 3, 21 and 7
// Product of all Subarrays = 30870

package main

import "fmt"

func main() {
	arr := [3]int{10, 3, 7}
	productSubarrays(arr)
}

func productSubarrays(arr [3]int) {
	product := 1
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			product *= arr[j]
		}
	}
	fmt.Println(product)
}
