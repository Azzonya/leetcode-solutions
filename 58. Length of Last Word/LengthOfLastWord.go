package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}
