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
	// Falls on the first day of the year

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

func MartinLutherKingJuniorDay(num int) []Holiday {
	// Falls on the third monday of January

	var holidays []Holiday
	d := BeginningOfMonth(PreviousYear(time.Now()))

	for len(holidays) < num {
		if d.Month() == time.January {
			holidays = append(
				holidays,
				Holiday{
					"Martin Luther King Jr. Day",
					WeeksSince(NextDayOfWeek(d, time.Monday), 2),
					"MLK",
				},
			)
			d = NextYear(d)
		} else {
			d = NextMonth(d)
		}
	}
	return holidays
}

func PresidentsDay(num int) []Holiday {
	// Falls on the third monday of February

	var holidays []Holiday
	d := BeginningOfMonth(PreviousYear(time.Now()))

	for len(holidays) < num {
		if d.Month() == time.February {
			holidays = append(
				holidays,
				Holiday{
					"Presidents Day",
					WeeksSince(NextDayOfWeek(d, time.Monday), 2),
					"PRS",
				},
			)
			d = NextYear(d)
		} else {
			d = NextMonth(d)
		}
	}
	return holidays
}

func IndependenceDay(num int) []Holiday {
	var holidays []Holiday
	d := PreviousMonthOfYear(PreviousYear(time.Now()), time.July)
	for len(holidays) < num {
		july4th := d.AddDate(0, 0, 3)
		publicHoliday := july4th
		if july4th.Weekday() == time.Saturday {
			publicHoliday = PreviousDay(july4th)
		} else if july4th.Weekday() == time.Sunday {
			publicHoliday = NextDay(july4th)
		}
		holidays = append(
			holidays,
			Holiday{
				"Independence Day",
				publicHoliday,
				"IND",
			},
		)
		d = NextYear(d)
	}
	return holidays
}
