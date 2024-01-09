package main

import "fmt"

func mySqrt(x int) int {

	startIncl := 0
	endIncl := x
	guess := (startIncl + endIncl) / 2

	for startIncl <= endIncl {
		guessSqr := guess * guess
		if guessSqr == x {
			return guess
		} else if guessSqr > x {
			endIncl = guess - 1
		} else {
			startIncl = guess + 1
		}
		guess = (startIncl + endIncl) / 2
	}

	return guess
}

func main() {
	result := mySqrt(27)
	fmt.Println(result)
}
