package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for range t {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		fields := strings.Fields(scanner.Text())

		a := make([]int, n)
		for i := range n {
			a[i], _ = strconv.Atoi(fields[i])
		}

		curMin := a[0]
		curLength := 0
		res := make([]int, 0)

		for i := range n {
			if min(curMin, a[i]) < curLength+1 {
				res = append(res, curLength)
				curLength = 1
				curMin = a[i]
			} else {
				curMin = min(curMin, a[i])
				curLength++
			}
		}

		res = append(res, curLength)

		resStr := make([]string, len(res))
		for i := range res {
			resStr[i] = strconv.Itoa(res[i])
		}

		writer.WriteString(strconv.Itoa(len(res)) + "\n")
		writer.WriteString(strings.Join(resStr, " ") + "\n")
	}
}
