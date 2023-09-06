package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 6, 4}
	fmt.Println(containsDuplicate(prices))
}

func containsDuplicate(nums []int) bool {
	valuesCounter := make(map[int]int)

	for _, v := range nums {
		valuesCounter[v]++
		if valuesCounter[v] > 1 {
			return true
		}
	}

	return false
}
