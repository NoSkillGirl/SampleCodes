// Given a binary string 'str' of length 'N' and an integer K,
// the task is to find the minimum number of steps required
// to move from str[0] to str[N – 1] with the following moves:

// 1. From an index i, the only valid moves are i + 1, i + 2 and i + K
// 2. An index i can only be visited if str[i] = ‘1’
// 3. Print -1 if it cant be reached.

// Examples
// Input: str = “101000011”, K = 5
// Output: 3

// Input: str = “1100000100111”, K = 6
// Output: -1

// Input: str = “10101010101111010101”, K = 4
// Output: 6

package main

import (
	"fmt"
	"time"
)

func main() {
	output := countSteps("10101010101111010101", 20, 4)
	fmt.Println("output:", output)
}

func countSteps(str string, n int, k int) int {
	//string ends with zero
	if string(str[n-1]) == "0" {
		return 0
	}
	//string have only one element
	if n == 1 {
		return 0
	}
	// string have exactly k elements
	if n == k-1 {
		return 1
	}

	count := 0
	l := n
	i := 0
	var oneFound bool
	for l > 1 {
		oneFound = false
		if k > 2 {
			if l > k {
				if string(str[i+k]) == "1" {
					count++
					i = i + k
					l = l - k
					oneFound = true
					continue
				}
			}
		}
		if l > 2 {
			if string(str[i+2]) == "1" {
				count++
				i = i + 2
				l = l - 2
				oneFound = true
				continue
			}
		}
		if l > 1 {
			if string(str[i+1]) == "1" {
				count++
				i++
				l--
				oneFound = true
				continue
			}
		}
		if oneFound == false {
			return -1
		}
	}
	return count
}
