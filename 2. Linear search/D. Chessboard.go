package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var chess [10][10]int

	for range n {
		var x, y int
		fmt.Scan(&x, &y)
		chess[x][y] = 1
	}

	p := 0

	for i := 1; i < 9; i++ {
		for j := 1; j < 9; j++ {
			if chess[i][j] == 1 {
				sum := 0
				sum += chess[i-1][j]
				sum += chess[i][j-1]
				sum += chess[i+1][j]
				sum += chess[i][j+1]
				p += 4 - sum
			}
		}
	}

	fmt.Println(p)
}
