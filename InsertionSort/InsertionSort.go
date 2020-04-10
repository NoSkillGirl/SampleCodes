// Insertion sort is a simple sorting algorithm that works the way we sort playing cards in our hands.
// Algorithm
// // Sort an arr[] of size n
// insertionSort(arr, n)
// Loop from i = 1 to n-1.
// ……a) Pick element arr[i] and insert it into sorted sequence arr[0…i-1]

package main

import (
	"fmt"
)

func main() {
	arr := [5]int{64, 25, 12, 22, 11}
	insertionSort(arr)
}

func insertionSort(arr [5]int) {
	fmt.Println("array: ", arr)
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	fmt.Println("sorted array: ", arr)
}
