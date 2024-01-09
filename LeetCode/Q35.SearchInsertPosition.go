package main

import "fmt"

func tryBinarySearch(nums []int, target int) (bool, int) {

	if nums == nil {
		panic("undefined slice")
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return true, middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return false, left
}

func searchInsert(nums []int, target int) int {
	_, result := tryBinarySearch(nums, target)
	return result
}

func main() {
	arr := []int{1, 3, 5, 6}
	target := -1
	result := searchInsert(arr, target)
	fmt.Println(result)
}
