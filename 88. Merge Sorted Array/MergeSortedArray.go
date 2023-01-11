package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1}
	nums2 := []int{}
	merge(nums1, 1, nums2, 0)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m] = nums2[i]
		m++
	}
	sort.Ints(nums1)
}
