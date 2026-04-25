package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	cost := make([]int, n)
	for i := range cost {
		fmt.Scan(&cost[i])
	}

	res := 0
	for i := range n {
		for j := range k {
			if i+j+1 >= n {
				break
			}
			res = max(cost[i+j+1]-cost[i], res)
		}
	}

	fmt.Println(res)
}
