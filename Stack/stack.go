package main

import "fmt"

/*func main() {
	var arr [3]string
	endIndex := 0
	var err string
	var lastEle string
	var arrEmpty bool
	var arrFull bool

	arrEmpty = isEmpty(endIndex)
	fmt.Println(arrEmpty)

	lastEle, err = peek(arr, endIndex)
	fmt.Println(arr, lastEle, err)

	arr, endIndex, err = push("dsfsf", arr, endIndex)
	fmt.Println(arr, err, endIndex)
	arr, endIndex, err = push("dsfgbhgb", arr, endIndex)
	fmt.Println(arr, err, endIndex)

	arrEmpty = isEmpty(endIndex)
	fmt.Println(arrEmpty)

	arr, endIndex, err = push("gjbhkj", arr, endIndex)
	fmt.Println(arr, err, endIndex)
	arr, endIndex, err = push("d", arr, endIndex)
	fmt.Println(arr, err, endIndex)

	arrFull = isFull(endIndex)
	fmt.Println(arrFull)

	arr, endIndex, err = pop(arr, endIndex)
	fmt.Println(arr, err, endIndex)

	lastEle, err = peek(arr, endIndex)
	fmt.Println(arr, lastEle, err)

	arr, endIndex, err = pop(arr, endIndex)
	fmt.Println(arr, err, endIndex)

	arrFull = isFull(endIndex)
	fmt.Println(arrFull)

	arr, endIndex, err = pop(arr, endIndex)
	fmt.Println(arr, err, endIndex)
	arr, endIndex, err = pop(arr, endIndex)
	fmt.Println(arr, err, endIndex)
}

func push(s string, inArr [3]string, endIndex int) (finalArray [3]string, returnEndIndex int, error string) {
	var arrFull bool
	arrFull = isFull(endIndex)
	fmt.Println(arrFull)

	if arrFull == false {
		inArr[endIndex] = s
		returnEndIndex = endIndex + 1
		finalArray = inArr
		return finalArray, returnEndIndex, "nil"
	}
	return inArr, endIndex, "err"
}

func pop(inArr [3]string, endIndex int) (finalArray [3]string, returnEndIndex int, error string) {
	var arrEmpty bool
	arrEmpty = isEmpty(endIndex)
	fmt.Println(arrEmpty)
	if (endIndex > 0) && (arrEmpty == false) {
		inArr[endIndex-1] = ""
		returnEndIndex = endIndex - 1
		finalArray = inArr
		return finalArray, returnEndIndex, "nil"
	}
	return inArr, endIndex, "err"
}

func peek(inArr [3]string, endIndex int) (lastElement string, error string) {
	if endIndex > 0 {
		lastElement = inArr[endIndex-1]
		return lastElement, "nil"
	}
	return "", "err"
}

func isEmpty(endIndex int) (emptyStatus bool) {
	if endIndex == 0 {
		return true
	}
	return false
}

func isFull(endIndex int) (fullStatus bool) {
	if endIndex >= 2 {
		return true
	}
	return false
}
*/

func main() {
	var arr [3]string
	endIndex := 0
	var err string
	var lastEle string
	endIndex, err = push("rttfgyuuh", &arr, endIndex)
	fmt.Println(arr, endIndex, err)
	lastEle, err = peek(&arr, endIndex)
	fmt.Println(arr, lastEle, err)
	endIndex, err = pop(&arr, endIndex)
	fmt.Println(arr, endIndex, err)
}

func push(s string, inArr *[3]string, endIndex int) (returnEndIndex int, error string) {
	var arrFull bool
	arrFull = isFull(endIndex)
	//fmt.Println(arrFull)
	if arrFull == false {
		(*inArr)[endIndex] = s
		returnEndIndex = endIndex + 1
		return returnEndIndex, "nil"
	}
	return endIndex, "err"
}

func pop(inArr *[3]string, endIndex int) (returnEndIndex int, error string) {
	var arrEmpty bool
	arrEmpty = isEmpty(endIndex)
	//fmt.Println(arrEmpty)
	if (endIndex > 0) && (arrEmpty == false) {
		(*inArr)[endIndex-1] = ""
		returnEndIndex = endIndex - 1
		return returnEndIndex, "nil"
	}
	return endIndex, "err"
}

func isFull(endIndex int) (fullStatus bool) {
	if endIndex >= 2 {
		return true
	}
	return false
}

func isEmpty(endIndex int) (emptyStatus bool) {
	if endIndex == 0 {
		return true
	}
	return false
}

func peek(inArr *[3]string, endIndex int) (lastElement string, error string) {
	if endIndex > 0 {
		lastElement = (*inArr)[endIndex-1]
		return lastElement, "nil"
	}
	return "", "err"
}
