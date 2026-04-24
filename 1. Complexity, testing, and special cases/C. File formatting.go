package main

import "fmt"

func main() {
	var n, res int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)

		tabs := a / 4
		left := a % 4

		if left <= 2 {
			res += tabs + left
		} else {
			res += tabs + 2
		}
	}

	fmt.Println(res)
}
