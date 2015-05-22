package main

import (
	"time"
)

const oneMonth = time.Duration(24) * time.Hour

type HolidayInstance struct {
	Description string
	Date        time.Time
	Code        string
}

type HolidayGenerator func(num int) []HolidayInstance

func MemorialDay(num int) []HolidayInstance {
	// Memorial Day falls on the last monday in May

	var instances []HolidayInstance
	d := BeginningOfMonth(time.Now().AddDate(-1, 0, 0))

	for len(instances) < num {
		if d.Month() == time.May {
			instances = append(
				instances,
				HolidayInstance{
					"Memorial Day",
					PreviousDayOfWeek(EndOfMonth(d), time.Monday),
					"MEM",
				},
			)
			d = NextYear(d)
		} else {
			d = NextMonth(d)
		}
	}
	return instances
}
