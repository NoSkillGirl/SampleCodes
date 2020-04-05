// Find maximum value of Sum( i*arr[i]) with only rotations on given array allowed
// Given an array, only rotation operation is allowed on array.
// We can rotate the array as many times as we want. Return the maximum possbile of summation of i*arr[i].

// Time Complexity : O(n)
// Auxiliary Space : O(1)

// Input: arr[] = {1, 20, 2, 10}
// Output: 72
// We can 72 by rotating array twice.
// {2, 10, 1, 20}
// 20*3 + 1*2 + 10*1 + 2*0 = 72

// Rj - Rj-1 = arrSum - n * arr[n-j]

package main

import "fmt"

func main() {
	arr := [4]int{1, 20, 2, 10}
	n := len(arr)
	arrSum := 0
	currVal := 0
	for i := 0; i < n; i++ {
		arrSum = arrSum + arr[i]
		currVal = currVal + i*arr[i]
	}

	maxVal := currVal
	for j := 1; j < n-1; j++ {
		currVal = currVal + arrSum - n*arr[n-j]
		if currVal > maxVal {
			maxVal = currVal
		}
	}

	fmt.Println("max value: ", maxVal)
}
