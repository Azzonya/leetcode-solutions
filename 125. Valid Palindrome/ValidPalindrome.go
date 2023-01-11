package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("Azamat"))
}

func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	processedString := strings.ToLower(reg.ReplaceAllString(s, ""))

	r := []rune(processedString)
	comparingIndex := len(r) - 1
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[comparingIndex] {
			return false
		}

		comparingIndex--
		//for j := len(r) - 1; j > len(r)/2; j-- {
		//	if r[i] != r[j] {
		//		return false
		//	}
		//}
	}
	return true
}
