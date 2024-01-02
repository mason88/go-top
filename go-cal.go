package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	// Get current year, month
	now := time.Now()
	year, month := now.Year(), now.Month()

	// Print heading with month, year
	heading := fmt.Sprintf("\n\t%s %d\n", now.Format("January"), year)
	fmt.Println(heading)

	// Print day headings
	days := []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	var daysRow strings.Builder
	for _, d := range days {
		daysRow.WriteString("\t" + d + "\t")
	}
	fmt.Println(daysRow.String())

	// Print days
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1).Day()

	empty := "\t\t"
	var week strings.Builder
	weekRow := 0
	for day := 1; day <= lastDay; day++ {
		if firstDay.Weekday() == time.Sunday && day == 1 {
			weekRow = 1
		}

		if weekRow == 0 {
			week.WriteString(empty)
		} else {
			week.WriteString("\t" + fmt.Sprint(day) + "\t")
		}

		if firstDay.Weekday() == time.Saturday {
			fmt.Println(week.String())
			week.Reset()
			weekRow = 0
		} else {
			weekRow++
		}

		firstDay = firstDay.AddDate(0, 0, 1)
	}

	if week.Len() > 0 {
		fmt.Println(week.String())
	}
}

