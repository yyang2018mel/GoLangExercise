package main

import "fmt"

func moveTarget(nums []int, target int) {
	fast := 0
	slow := fast

	for fast < len(nums) {
		if nums[fast] == target {
			fast = fast + 1

		} else {
			tmp := nums[slow]
			nums[slow] = nums[fast]
			nums[fast] = tmp

			fast = fast + 1
			slow = slow + 1
		}
	}
}

func moveZeroes(nums []int) {
	moveTarget(nums, 0)
}

func main() {
	arr := []int{3, 3, 0, 1, 2, 0, 3, 0}
	moveZeroes(arr)
	fmt.Println(arr)
}
