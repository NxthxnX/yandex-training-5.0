package main

import (
	"fmt"
	"slices"
)

func main() {
	var n, a, b, k int

	fmt.Scan(&n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	fmt.Scan(&a, &b, &k)

	allButFirst := make([]int, n-1)
	copy(allButFirst, nums[1:])
	slices.Reverse(allButFirst)

	res := max(findMax(a, b, k, nums), findMax(a, b, k, append([]int{nums[0]}, allButFirst...)))

	fmt.Println(res)
}

func findMax(a, b, k int, nums []int) int {
	n := len(nums)
	seen := make([]bool, n)
	landA := (a - 1) / k
	landB := (b - 1) / k

	maxVal := 0

	for i := landA; i <= landB; i++ {
		if seen[i%n] {
			break
		}
		seen[i%n] = true
		maxVal = max(maxVal, nums[i%n])
	}

	return maxVal
}
