package main

import (
	"time"
)

const oneMonth = time.Duration(24) * time.Hour

type Holiday struct {
	Description string
	Date        time.Time
	Code        string
}

type HolidayGenerator func(num int) []Holiday

func NewYearsDay(num int) []Holiday {
	var holidays []Holiday
	d := BeginningOfYear(time.Now())
	for len(holidays) < num {
		holidays = append(
			holidays,
			Holiday{
				"New Year's Day",
				d,
				"NEW",
			},
		)
		d = NextYear(d)
	}
	return holidays
}

func MemorialDay(num int) []Holiday {
	// Memorial Day falls on the last monday in May

	var holidays []Holiday
	d := BeginningOfMonth(PreviousYear(time.Now()))

	for len(holidays) < num {
		if d.Month() == time.May {
			holidays = append(
				holidays,
				Holiday{
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
	return holidays
}
