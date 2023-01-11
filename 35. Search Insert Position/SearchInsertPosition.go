package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))
}

func searchInsert(nums []int, target int) int {
	output := 0
	for i, r := range nums {
		if r == target {
			return i
		}
		if r < target {
			output++
		}
	}
	return output
}
