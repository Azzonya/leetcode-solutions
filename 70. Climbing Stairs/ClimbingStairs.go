package main

import "fmt"

func main() {
	fmt.Println(climbStairs(6))
}

func climbStairs(n int) int {
	//if n == 1 || n == 2 || n == 0 {
	//	return n
	//}
	// s := math.Sqrt(float64(5))
	// fib := math.Pow((1+s)/2, float64(n)+1) - math.Pow((1-s)/2, float64(n)+1)

	// return int(fib / s)
	if n <= 1 {
		return n
	}
	pre, cur := 1, 1
	for i := 1; i < n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}
