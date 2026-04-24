package main

import "fmt"

func main() {
	var p, v, q, m int
	fmt.Scan(&p, &v, &q, &m)

	l1, r1 := p-v, p+v
	l2, r2 := q-m, q+m

	len1 := r1 - l1 + 1
	len2 := r2 - l2 + 1

	overlapL := max(l1, l2)
	overlapR := min(r1, r2)

	overlap := max(overlapR-overlapL+1, 0)

	fmt.Println(len1 + len2 - overlap)
}
