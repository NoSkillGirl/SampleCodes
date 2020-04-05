package main

import "fmt"

func main() {
	var err string

	arr := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(arr)

	arr, err = changeElement(arr, 4, 45)
	fmt.Println(arr, err)
	arr, err = changeElement(arr, 8, 45)
	fmt.Println(arr, err)
	arr, err = changeElement(arr, 0, 45)
	fmt.Println(arr, err)

}

func changeElement(inArr [6]int, pos int, replacement int) (finalArray [6]int, error string) {
	if pos < 7 || pos > 0 {
		inArr[pos-1] = replacement
		return inArr, "nil"
	}
	return inArr, "error"
}
