package main

import (
	"fmt"
	"math"
)

func main() {
	var l, x1, v1, x2, v2 int
	fmt.Scan(&l, &x1, &v1, &x2, &v2)

	ans := math.MaxFloat64
	ans = min(ans, calcTime(x1, v1, x2, v2, l))
	ans = min(ans, calcTime(x1, v1, l-x2, -v2, l))

	if ans == math.MaxFloat64 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Printf("%.10f\n", ans)
	}
}

func calcTime(x1, v1, x2, v2, l int) float64 {
	dx := ((x2-x1)%l + l) % l
	dv := v1 - v2

	if dx == 0 {
		return 0
	}

	if dv < 0 {
		dv = -dv
		dx = l - dx
	}

	if dv > 0 {
		return float64(dx) / float64(dv)
	}

	return math.MaxFloat64
}
