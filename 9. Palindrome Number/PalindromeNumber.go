package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	chars := []rune(xStr)
	elementsQuantity := len(chars) - 1
	for i := 0; i < len(chars); i++ {
		//compareElement := chars[elementsQuantity]
		if chars[i] != chars[elementsQuantity] {
			return false
		}
		elementsQuantity -= 1
		//for j := len(chars); j > 0; j-- {
		//	fmt.Println(i, j)
		//}
	}
	return true
}
