package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	nextInt := func() int {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		return val
	}

	n := nextInt()

	a := make([]int, n)
	b := make([]int, n)

	for i := range n {
		a[i] = nextInt()
		b[i] = nextInt()
	}

	var (
		good, bad               []int
		sumGood                 int
		maxGoodBIdx, maxBadAIdx int = -1, -1
	)

	for i := range n {
		if a[i] >= b[i] {
			good = append(good, i)
			sumGood += a[i] - b[i]
			if maxGoodBIdx == -1 || b[i] > b[maxGoodBIdx] {
				maxGoodBIdx = i
			}
		} else {
			bad = append(bad, i)
			if maxBadAIdx == -1 || a[i] > a[maxBadAIdx] {
				maxBadAIdx = i
			}
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	writeInt := func(val int) {
		writer.WriteString(strconv.Itoa(val))
	}
	writeSpace := func() {
		writer.WriteByte(' ')
	}
	writeNewline := func() {
		writer.WriteByte('\n')
	}

	if maxGoodBIdx == -1 {
		writeInt(a[maxBadAIdx])
		writeNewline()

		writeInt(maxBadAIdx + 1)
		writeSpace()

		for i := range n {
			if i != maxBadAIdx {
				writeInt(i + 1)
				writeSpace()
			}
		}
	} else if maxBadAIdx == -1 || b[maxGoodBIdx] > a[maxBadAIdx] {
		writeInt(sumGood + b[maxGoodBIdx])
		writeNewline()

		for _, i := range good {
			if i != maxGoodBIdx {
				writeInt(i + 1)
				writeSpace()
			}
		}
		writeInt(maxGoodBIdx + 1)
		writeSpace()

		for _, i := range bad {
			writeInt(i + 1)
			writeSpace()
		}
	} else {
		writeInt(sumGood + a[maxBadAIdx])
		writeNewline()

		for _, i := range good {
			writeInt(i + 1)
			writeSpace()
		}
		writeInt(maxBadAIdx + 1)
		writeSpace()

		for _, i := range bad {
			if i != maxBadAIdx {
				writeInt(i + 1)
				writeSpace()
			}
		}
	}
}
