package main

import "fmt"

func isPerfectSquare(num int) bool {

	start := 0
	end := num

	for start <= end {
		guess := (start + end) / 2
		guessSqr := guess * guess

		if guessSqr == num {
			return true
		} else if guessSqr > num {
			end = guess - 1
		} else {
			start = guess + 1
		}
	}
	return false
}

func main() {
	is := isPerfectSquare(25)
	fmt.Println(is)
}
