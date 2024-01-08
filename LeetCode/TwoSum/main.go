package main

import "fmt"

func makeValueToIndexMapping(nums []int) map[int]int {
	mapping := make(map[int]int)
	for idx, val := range nums {
		mapping[val] = idx
	}
	return mapping
}

func twoSum(nums []int, target int) []int {
	valToIdxMapping := makeValueToIndexMapping(nums)
	for idx, num := range nums {
		subTarget := target - num
		if sndIdx, found := valToIdxMapping[subTarget]; found && sndIdx != idx {
			return []int{idx, sndIdx}
		}
	}
	return make([]int, 2)
}

func main() {
	arr := []int{3, 0, 3, 4}
	target := 6
	result := twoSum(arr, target)
	fmt.Println(result)
}
