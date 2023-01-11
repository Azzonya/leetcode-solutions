package main

func main() {
	nums := []int{2, 7, 11, 15}
	twoSum(nums, 9)
}

func twoSum(nums []int, target int) []int {
	var output = []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				output = append(output, i)
				output = append(output, j)
			}
		}
	}
	return output
}
