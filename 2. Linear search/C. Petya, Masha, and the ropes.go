package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	ropes := make([]int, n)
	for i := range ropes {
		fmt.Scan(&ropes[i])
	}

	maxRope := ropes[0]
	sum := 0
	for _, r := range ropes {
		if r > maxRope {
			maxRope = r
		}
		sum += r
	}

	if sum-maxRope < maxRope {
		fmt.Println(maxRope - (sum - maxRope))
	} else {
		fmt.Println(sum)
	}
}
