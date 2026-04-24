package main

import "fmt"

func main() {
	var firstTeam1, secondTeam1, firstTeam2, secondTeam2, place int
	fmt.Scanf("%d:%d\n%d:%d\n%d", &firstTeam1, &secondTeam1, &firstTeam2, &secondTeam2, &place)

	sumFirst := firstTeam1 + firstTeam2
	sumSecond := secondTeam1 + secondTeam2

	res := sumSecond - sumFirst

	if res >= 0 {
		switch place {
		case 1:
			if secondTeam1 >= firstTeam2+res {
				res++
			}
		case 2:
			if secondTeam2 >= firstTeam1 {
				res++
			}
		}
	} else {
		res = 0
	}

	fmt.Println(res)
}
