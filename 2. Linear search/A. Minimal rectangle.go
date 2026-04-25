package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	fmt.Scan(&k)

	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt

	for range k {
		var x, y int
		fmt.Scan(&x, &y)

		minX = min(minX, x)
		maxX = max(maxX, x)
		minY = min(minY, y)
		maxY = max(maxY, y)
	}

	fmt.Println(minX, minY, maxX, maxY)
}
