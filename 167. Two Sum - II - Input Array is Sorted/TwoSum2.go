package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(numbers []int, target int) []int {
	var output = []int{}
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				output = append(output, i+1)
				output = append(output, j+1)
			}
		}
	}
	return output

}
