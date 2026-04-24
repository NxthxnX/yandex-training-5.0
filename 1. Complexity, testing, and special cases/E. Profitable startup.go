package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		n, k, d int
		flag    bool
		res     string
	)

	fmt.Scan(&n, &k, &d)

	modulo := n % k

	for digit := range 10 {
		if (modulo*10+digit)%k == 0 {
			flag = true
			res = strconv.Itoa(n) + strconv.Itoa(digit) + strings.Repeat("0", d-1)
			break
		}
	}

	if flag {
		fmt.Println(res)
	} else {
		fmt.Println(-1)
	}
}
