// Print last k digits of a^b (a raised to power b)
// Given positive integers k, a and b we need to print last k digits of a^b ie.. pow(a, b).

// Input Constraint:
// k <= 9, a <= 10^6, b<= 10^6
// Examples:

// Input : a = 11, b = 3, k = 2
// Output : 31
// Explanation : a^b = 11^3 = 1331, hence
// last two digits are 31

package main

import "fmt"

func main() {
	printLastKDigits(11, 3, 2)
	printLastKDigits(10, 10000, 5)
}

func printLastKDigits(a, b, k int) {
	fmt.Printf("\nlast %d digits of %d^%d = ", k, a, b)

	temp := 1
	for i := 1; i <= k; i++ {
		temp = temp * 10
	}

	temp = power(a, b, temp)

	for i := 0; i < k-numberOfDigits(temp); i++ {
		fmt.Print(0)
	}

	if temp != 0 {
		fmt.Println(temp)
	}
}

func power(x, y, p int) int {
	res := 1
	x = x % p
	for y > 0 {
		if y%2 != 0 {
			res = (res * x) % p
		}
		y = y / 2
		x = (x * x) % p
	}
	return res
}

func numberOfDigits(x int) int {
	i := 0
	for x > 0 {
		x /= 10
		i++
	}
	return i
}
