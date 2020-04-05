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
func Rotate(inArr *[5]int, noOfElement int) {
	if noOfElement < 1 || noOfElement > 4 {
		fmt.Println("Enter a number between 1-4")
	} else {
		k := noOfElement + 1
		var s []int = (*inArr)[1:k]

		for i := 0; i < (5 - noOfElement); i++ {
			(*inArr)[i] = (*inArr)[i+noOfElement]
		}
		for j := 0; j < noOfElement; j++ {
			l := j + 1
			(*inArr)[5-noOfElement+j] = s[j:l][0]
		}
	}
}
