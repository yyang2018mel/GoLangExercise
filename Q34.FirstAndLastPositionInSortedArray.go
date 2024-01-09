package main

import "fmt"

func findLeftBorder(nums []int, target int) int {
	leftBorder := -1

	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			leftBorder = middle
			right = middle - 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return leftBorder
}

func findRightBorder(nums []int, target int) int {
	rightBorder := -1

	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			rightBorder = middle
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return rightBorder
}

func searchRange(nums []int, target int) []int {
	left := findLeftBorder(nums, target)
	right := findRightBorder(nums, target)
	return []int{left, right}
}

func main() {
	arr := []int{1, 2, 3, 4, 4, 4, 8, 9, 10}
	target := 0
	left := findLeftBorder(arr, target)
	right := findRightBorder(arr, target)
	fmt.Println(left, right)
}
