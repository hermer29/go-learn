package leetcode

import (
	"slices"
)

func twoSum(nums []int, target int) []int {
	values := make(map[int]int)

	for i := range nums {
		subtruction := target - nums[i]
		index, exists := values[subtruction]
		if exists {
			return []int{i, index}
		} else {
			values[nums[i]] = i
		}
	}
	return nil
}

func threeSum(nums []int) [][]int {
	results := [][]int{}
	slices.Sort(nums)
	for j := range nums {
		values := make(map[int]int)

		for i := range nums[j+1 : len(nums)-1] {
			subtruction := nums[j] + nums[i+j+1]
			index, exists := values[-subtruction]
			if exists {
				results = append(results, []int{nums[i], nums[index], nums[j]})
			} else {
				values[nums[i]] = i
			}
		}
	}

	return results
}
