package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
}

func removeElement(nums []int, val int) int {
	numsLen := len(nums)
	i := 0
	for i < numsLen {
		if nums[i] == val {
			numsLen--
			nums[i] = nums[numsLen]
		} else {
			i++
		}
	}

	return numsLen
}
