package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	counter := 0
	for i, r := range nums {
		if i < len(nums)-1 && r == nums[i+1] {
			continue
		}
		nums[counter] = r
		counter++
	}

	return counter
}
