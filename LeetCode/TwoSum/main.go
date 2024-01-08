package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for idx, num := range nums {
		subTarget := target - num
		if sndIdx, found := mapping[subTarget]; found {
			return []int{idx, sndIdx}
		}
		mapping[num] = idx
	}
	return make([]int, 2)
}

func main() {
	arr := []int{3, 0, 3, 4}
	target := 6
	result := twoSum(arr, target)
	fmt.Println(result)
}
