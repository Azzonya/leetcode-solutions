package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
func search(nums []int, target int) int {
	for i, r := range nums {
		if r == target {
			return i
		}
	}
	return -1
}
