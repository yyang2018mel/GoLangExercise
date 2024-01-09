package main

import "fmt"

func removeElement(nums []int, val int) int {

	current := 0
	end := len(nums) - 1

	for current <= end {
		if nums[current] == val {
			valEnd := nums[end]
			nums[end] = nums[current]
			nums[current] = valEnd
			end = end - 1
		} else {
			current = current + 1
		}
	}

	count := end + 1
	return count
}

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	result := removeElement(arr, val)
	fmt.Println(result)
}
