package main

import (
	"fmt"
	"slices"
)

type Date struct {
	Day   int
	Month int
}

func main() {
	var (
		n, year      int
		firstDayName string
		isLeapYear   int
	)

	fmt.Scan(&n, &year)

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		isLeapYear = 1
	}

	var monthNames = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	var daysInMonth = []int{31, 28 + isLeapYear, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var dayNames = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	holidays := make(map[Date]struct{}, n)
	for i := 0; i < n; i++ {
		var (
			day   int
			month string
		)
		fmt.Scan(&day, &month)

		holidays[Date{day, slices.Index(monthNames, month)}] = struct{}{}
	}

	fmt.Scan(&firstDayName)

	daysCount := make([]int, len(dayNames))
	holidaysCount := make([]int, len(dayNames))

	curDayOfMonth := 1
	curMonth := 0
	curDayOfWeek := slices.Index(dayNames, firstDayName)

	for range 365 + isLeapYear {
		if _, ok := holidays[Date{curDayOfMonth, curMonth}]; ok {
			holidaysCount[curDayOfWeek]++
		}
		daysCount[curDayOfWeek]++

		if curDayOfMonth == daysInMonth[curMonth] {
			curDayOfMonth = 1
			curMonth++
		} else {
			curDayOfMonth++
		}
		curDayOfWeek = (curDayOfWeek + 1) % len(dayNames)
	}

	minDays := 54
	maxDays := -1

	var bestDayOfWeek, worstDayOfWeek int

	for weekday := range dayNames {
		workDays := daysCount[weekday] - holidaysCount[weekday]

		if workDays < minDays {
			minDays = workDays
			worstDayOfWeek = weekday
		}
		if workDays > maxDays {
			maxDays = workDays
			bestDayOfWeek = weekday
		}
	}

	fmt.Println(dayNames[bestDayOfWeek], dayNames[worstDayOfWeek])
}
