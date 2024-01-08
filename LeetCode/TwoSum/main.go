package main

import "fmt"

func tryFindTarget(nums []int, length int, target int) (bool, int) {
	for i := 0; i < length; i++ {
		if nums[i] == target {
			return true, i
		}
	}
	return false, -1
}

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	for idx := 0; idx < numsLen; idx++ {
		subTarget := target - nums[idx]
		found, offset := tryFindTarget(nums[idx+1:], numsLen-idx-1, subTarget)
		if found {
			return []int{idx, idx + 1 + offset}
		}
	}
	return make([]int, 2)
}

func main() {
	arr := []int{1, 2, 3}
	target := 4
	result := twoSum(arr, target)
	fmt.Println(result)
}
