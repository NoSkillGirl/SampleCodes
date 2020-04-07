package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 3, 5, 7, 9}
	printEle(arr, 1)
	printEle(arr, 2)
	printEle(arr, 3)
}

func printEle(arr [5]int, k int) {
	fmt.Println("Rotated Array by ", k, ":")
	n := len(arr)
	mod := k % n
	for i := 0; i < n; i++ {
		fmt.Println(arr[(mod+i)%n])
	}
}
