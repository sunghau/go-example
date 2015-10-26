package main

import (
	"fmt"
	"time"
)

const (
    // See http://golang.org/pkg/time/#Parse
    timeFormat = "2006-01-02 15:04 MST"
)

func showMonthCalendar(t time.Time) {
	year, month, _ := t.Date()

	fmt.Printf("\t\t%v\t\t%v\n", month, year)
	fmt.Println("Sun\tMon\tTue\tWed\tThu\tFri\tSat")

	cal := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := cal.AddDate(0, 1, 0)
	if end.Weekday() != 0 {
		end = end.AddDate(0, 0, (int)(6-end.Weekday()))
	} else {
		end = end.AddDate(0, 0, -1)
	}

	cal = cal.AddDate(0, 0, (int)(-cal.Weekday()))

	for !end.Before(cal) {
		if cal.Weekday() == 6 {
			fmt.Println(cal.Day())
		} else {
			fmt.Print(cal.Day(), "\t")
		}
		cal = cal.AddDate(0, 0, 1)
	}
}

func main() {
	showMonthCalendar(time.Now())
}
