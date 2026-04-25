package main

import (
	"fmt"
	"math"
	"slices"
)

type Point struct {
	X, Y int
}

func main() {
	var n int
	fmt.Scan(&n)

	points := make([]Point, n)
	for i := range n {
		fmt.Scan(&points[i].X, &points[i].Y)
	}

	slices.SortFunc(points, func(p1, p2 Point) int {
		if p1.X < p2.X {
			return -1
		}
		if p1.X > p2.X {
			return 1
		}
		return 0
	})

	ans := math.MaxInt

	for col := range n {
		sum := 0
		for i := range n {
			sum += intAbs(points[i].X-(i+1)) + intAbs(points[i].Y-(col+1))
		}
		ans = min(ans, sum)
	}

	fmt.Println(ans)
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
