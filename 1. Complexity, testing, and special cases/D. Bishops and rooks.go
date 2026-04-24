package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var chessDesk [10][10]rune
	for j := range 10 {
		chessDesk[0][j] = '!'
		chessDesk[9][j] = '!'
	}
	for i := range 10 {
		chessDesk[i][0] = '!'
		chessDesk[i][9] = '!'
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= 8; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j := 1; j <= 8; j++ {
			chessDesk[i][j] = rune(line[j-1])
		}
	}

	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if chessDesk[i][j] == 'B' || chessDesk[i][j] == 'R' {
				var di, dj [4]int
				switch chessDesk[i][j] {
				case 'B':
					di, dj = [4]int{1, 1, -1, -1}, [4]int{1, -1, -1, 1}
				case 'R':
					di, dj = [4]int{0, 0, 1, -1}, [4]int{1, -1, 0, 0}
				default:
					continue
				}

				for k := range 4 {
					ni, nj := i+di[k], j+dj[k]
					for chessDesk[ni][nj] == '*' || chessDesk[ni][nj] == '.' {
						chessDesk[ni][nj] = '.'
						ni, nj = ni+di[k], nj+dj[k]
					}
				}
			}
		}
	}

	var hit int

	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if chessDesk[i][j] == '.' || chessDesk[i][j] == 'B' || chessDesk[i][j] == 'R' {
				hit++
			}
		}
	}

	fmt.Println(8*8 - hit)
}
