package main

import "strings"

func main() {
	strStr("satbutsad", "sad")
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
