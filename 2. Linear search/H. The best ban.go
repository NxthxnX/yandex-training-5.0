package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	inputParams := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(inputParams[0])
	m, _ := strconv.Atoi(inputParams[1])

	pers := make([][]int, n)
	for i := range n {
		pers[i] = make([]int, m)
	}

	for i := range n {
		scanner.Scan()
		inputParams := strings.Fields(scanner.Text())
		for j := range m {
			pers[i][j], _ = strconv.Atoi(inputParams[j])
		}
	}

	_, mxR, mxC := findMaxWithDeleted(pers, -1, -1)

	_, _, delC := findMaxWithDeleted(pers, mxR, -1)
	_, delR, _ := findMaxWithDeleted(pers, -1, mxC)

	valDelC, _, _ := findMaxWithDeleted(pers, mxR, delC)
	valDelR, _, _ := findMaxWithDeleted(pers, delR, mxC)

	if valDelC < valDelR {
		writer.WriteString(strconv.Itoa(mxR+1) + " " + strconv.Itoa(delC+1) + "\n")
	} else {
		writer.WriteString(strconv.Itoa(delR+1) + " " + strconv.Itoa(mxC+1) + "\n")
	}
}

func findMaxWithDeleted(a [][]int, r, c int) (int, int, int) {
	n := len(a)
	m := len(a[0])
	maxVal := math.MinInt
	resR, resC := -1, -1

	for i := range n {
		for j := range m {
			if i != r && j != c {
				if maxVal < a[i][j] {
					maxVal = a[i][j]
					resR, resC = i, j
				}
			}
		}
	}

	return maxVal, resR, resC
}
