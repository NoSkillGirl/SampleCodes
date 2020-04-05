package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	Rotate(&arr, 2)
	fmt.Println(arr)
	Rotate(&arr, 3)
	fmt.Println(arr)
	Rotate(&arr, 0)
	Rotate(&arr, 5)
}

//Rotate function
func Rotate(inArr *[5]int, noOfElementsToRotate int) {
	if noOfElementsToRotate < 1 || noOfElementsToRotate > 4 {
		fmt.Println("Enter a number between 1-4")
	} else {
		for j := 0; j < noOfElementsToRotate; j++ {
			temp := (*inArr)[0]
			for i := 0; i < 4; i++ {
				(*inArr)[i] = (*inArr)[i+1]
			}
			(*inArr)[4] = temp
		}
	}
}
