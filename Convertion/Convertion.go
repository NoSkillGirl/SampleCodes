package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	s1 := "25"
	// string to int
	i, err := strconv.Atoi(s1)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(s1, i)

	// string to float
	s2 := "465.76"
	f, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(s2, f)

	// int to float
	a := 5
	b := float64(a)
	fmt.Printf("b is %f\n", b)

	// float to int
	var x float64 = 5.7
	var y int = int(x)
	fmt.Println(y)

	// Round float to 2 decimal places
	c := fmt.Sprintf("%.2f", 12.3456) // c == "12.35"
	fmt.Println("c: ", c)

	z := 12.3456
	fmt.Println(math.Floor(z*100) / 100) // 12.34 (round down)
	fmt.Println(math.Round(z*100) / 100) // 12.35 (round to nearest)
	fmt.Println(math.Ceil(z*100) / 100)  // 12.35 (round up)

}
