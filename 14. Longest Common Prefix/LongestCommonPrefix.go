package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	result := ""
	maxLen := maxPossiblePrefixLen(strs)
	for i := maxLen; i > 0; i-- {
		prefix := strs[0][0:i]
		equal := true
		for j := 1; j < len(strs); j++ {
			if prefix != strs[j][0:i] {
				equal = false
			}
		}
		if equal {
			return prefix
		}
		maxLen -= 1
	}

	return result
}

func maxPossiblePrefixLen(strs []string) int {
	if len(strs) == 0 {
		return 0
	}
	min := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if min > len(strs[i]) {
			min = len(strs[i])
		}
	}
	return min
}
