package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	output := x
	for output*output > x {
		output = (output + x/output) / 2
	}

	return output
}
