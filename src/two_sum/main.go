package main

import "fmt"

// <b>Two Sum<b>
// Given an array of integers, return indices of the two
// numbers such that they add up to a specific target

// You may assume that each input would have exactly one
// solution, and you may not use the same element twice.

// Example:
//		Given nums = [2, 7, 11, 15], target = 9
//
//		Because nums[0] + num[1] = 2 + 7 = 9
//		return [0, 1].

type indexes struct {
	x, y int
}

func main() {
	nums := []int{2, 3, 5, 11, 15, 7, 6, 4}
	result := twoSum(nums, 9)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if v, found := m[target-n]; found {
			return []int{v, i}
		}
		m[n] = i
	}
	return nil
}

func twoSumMulti(nums []int, target int) []indexes {
	m := make(map[int]int)
	result := make([]indexes, 0)
	for i, n := range nums {
		if v, found := m[target-n]; found {
			result = append(result, indexes{x: v, y: i})
		}
		m[n] = i
	}
	return result
}

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumBruteForceMulti(nums []int, target int) []indexes {
	result := make([]indexes, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, indexes{x: i, y: j})
			}
		}
	}
	return result
}
