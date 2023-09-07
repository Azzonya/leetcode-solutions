package _69__Majority_Element

func majorityElement(nums []int) int {
	valueCounter := make(map[int]int)

	for _, v := range nums {
		valueCounter[v]++
	}

	var maxValue int
	var maxKey int

	for k, v := range valueCounter {
		if v > maxValue {
			maxKey = k
			maxValue = v
		}
	}

	return maxKey
}
