package main

import "fmt"

func search_CloseOnBothEnds(nums []int, target int) int {

	if nums == nil || len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func search_LeftCloseRightOpen(nums []int, target int) int {

	if nums == nil || len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums)

	for left < right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return -1
}

func main() {
	arr := []int{2, 5, 6, 9, 10, 11, 12}
	result := search_LeftCloseRightOpen(arr, 9)
	fmt.Println(result)
}
