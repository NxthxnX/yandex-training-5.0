package main

import (
	"fmt"
)

func main() {
	var n int
	var res []rune
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	i := 0
	for i < n && nums[i]%2 == 0 {
		i++
		res = append(res, '+')
	}
	i++
	for i < n && (nums[i]%2 == 1 || nums[i]%2 == -1) {
		i++
		res = append(res, 'x')
	}
	if i < n {
		res = append(res, '+')
		i++
	}
	for i < n {
		res = append(res, 'x')
		i++
	}

	fmt.Println(string(res))
}
