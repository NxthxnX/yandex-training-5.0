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
	m, _ := strconv.Atoi(inputParams[0])
	n, _ := strconv.Atoi(inputParams[1])

	painting := make([][]rune, m)

	for i := range m {
		scanner.Scan()
		painting[i] = []rune(scanner.Text())[:n]
	}

	bools := [...]bool{true, false}
	isFound := false
	for _, isVertical := range bools {
		for line := 1; isVertical && line <= n-1 || !isVertical && line <= m-1; line++ {
			if isDividedIntoTwoRect(painting, line, isVertical) {
				paintRect(painting, line, isVertical)
				isFound = true
				break
			}
		}
		if isFound {
			break
		}
	}

	if isFound {
		writer.WriteString("YES\n")
		for _, row := range painting {
			writer.WriteString(string(row) + "\n")
		}
	} else {
		writer.WriteString("NO\n")
	}
}

func isDividedIntoTwoRect(painting [][]rune, line int, isVertical bool) bool {
	m := len(painting)
	n := len(painting[0])

	if isVertical {
		return isRect(painting, 0, m-1, 0, line-1) && isRect(painting, 0, m-1, line, n-1)
	} else {
		return isRect(painting, 0, line-1, 0, n-1) && isRect(painting, line, m-1, 0, n-1)
	}
}

func isRect(painting [][]rune, minI, maxI, minJ, maxJ int) bool {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt

	for i := minI; i <= maxI; i++ {
		for j := minJ; j <= maxJ; j++ {
			if painting[i][j] == '#' {
				minX = min(minX, i)
				maxX = max(maxX, i)
				minY = min(minY, j)
				maxY = max(maxY, j)
			}
		}
	}

	if minX == math.MaxInt {
		return false
	}

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if painting[i][j] == '.' {
				return false
			}
		}
	}

	return true
}

func paintRect(painting [][]rune, line int, isVertical bool) {
	m := len(painting)
	n := len(painting[0])

	for i := range m {
		for j := range n {
			if painting[i][j] == '#' {
				if isVertical {
					if j < line {
						painting[i][j] = 'a'
					} else {
						painting[i][j] = 'b'
					}
				} else {
					if i < line {
						painting[i][j] = 'a'
					} else {
						painting[i][j] = 'b'
					}
				}
			}
		}
	}
}
