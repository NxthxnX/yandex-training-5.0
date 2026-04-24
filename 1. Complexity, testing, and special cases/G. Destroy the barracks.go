package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, p int
	ans := math.MaxInt
	fmt.Scan(&x, &y, &p)

	for hp := range y + 1 {
		ans = min(ans, calcRound(x, y, p, hp))
	}

	if ans == math.MaxInt {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func calcRound(myUnits, townHp, enemyProd, targetHp int) int {
	rounds := 0
	enemyUnits := 0

	for townHp >= targetHp {
		if enemyUnits >= myUnits {
			return math.MaxInt
		}
		townHp -= myUnits - enemyUnits
		if townHp > 0 {
			enemyUnits = enemyProd
		}
		rounds++
	}

	for townHp > 0 {
		if myUnits <= 0 {
			return math.MaxInt
		}
		townHp -= myUnits
		if townHp > 0 {
			enemyUnits += enemyProd
		} else {
			enemyUnits += townHp
			townHp = 0
		}
		if enemyUnits > 0 {
			myUnits -= enemyUnits
		}
		rounds++
	}

	for enemyUnits > 0 {
		if myUnits <= 0 {
			return math.MaxInt
		}
		enemyUnits -= myUnits
		myUnits -= enemyUnits
		rounds++
	}

	return rounds
}
