//Write a function rotate(ar[], d, n) that rotates arr[] of size n by d elements.

// Time complexity : O(n)

package main

import (
	"fmt"
)

func main() {
	arr := [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	n := len(arr)
	d := 3
	t := (n / d) - 1

	//Reversal algorithm
	arrRotated1 := reverseAlgo(arr, 3)
	fmt.Println("Rotated array using reverse algo : ", arrRotated1)
	// Cyclically rotate
	arrRotated2 := reverseAlgo(arr, 3)
	fmt.Println("Rotated array using cyclic algo : ", arrRotated2)

	//A Juggling Algorithm
	//Time Complexity : O(n)
	// Auxiliary Space : O(1)
	gcd := calculateGCD(12, 3)
	fmt.Println("GCD: ", gcd)

	var lastEle int
	for i := 0; i < t; i++ {
		temp := arr[i]
		j := gcd + i
		for j < n {
			arr[j-gcd] = arr[j]
			lastEle = j
			j = j + gcd
		}
		arr[lastEle] = temp
	}
	fmt.Println("Array after rotation: ", arr)
}

func calculateGCD(temp1 int, temp2 int) int {
	var gcdnum int
	for i := 1; i <= temp1 && i <= temp2; i++ {
		if temp1%i == 0 && temp2%i == 0 {
			gcdnum = i
		}
	}
	return gcdnum
}

//Reversal algorithm for array rotation
// Time complexity : O(n)
func reverseAlgo(a [12]int, d int) [12]int {
	n := len(a)
	var b [12]int
	j := d - 1
	for i := 0; i < n; i++ {
		b[i] = a[j]
		if j == 0 {
			j = n - 1
		} else {
			j = j - 1
		}
	}

	j = n - 1
	for i := 0; i < n; i++ {
		a[i] = b[j]
		j = j - 1
	}
	return a
}

// Cyclically rotate an array by one
func cyclicallyRotation(a [12]int, d int) [12]int {
	n := len(a)
	for i := 0; i < d; i++ {
		temp := a[0]
		for j := 1; j < n; j++ {
			a[j-1] = a[j]
		}
		a[n-1] = temp
	}
	return a
}
