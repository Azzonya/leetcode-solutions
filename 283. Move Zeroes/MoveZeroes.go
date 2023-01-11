package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		curVal := nums[i]
		if curVal == 0 && i != len(nums)-1 {
			step := 1
			nextVal := nums[i+step]
			for nextVal == 0 {
				step++
				if i+step >= len(nums) {
					break
				}
				nextVal = nums[i+step]
			}
			if nextVal == 0 {
				continue
			}
			nums[i] = nextVal
			nums[i+step] = curVal
		}
	}
}
