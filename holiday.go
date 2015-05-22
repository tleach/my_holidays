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

type ByDate []Holiday

func (b ByDate) Len() int {
	return len(b)
}

func (b ByDate) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByDate) Less(i, j int) bool {
	return b[j].Date.After(b[i].Date)
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

func LaborDay(num int) []Holiday {
	// First monday in September
	var holidays []Holiday
	d := PreviousMonthOfYear(PreviousYear(time.Now()), time.September)

	for len(holidays) < num {
		holidays = append(
			holidays,
			Holiday{
				"Labor Day",
				NextDayOfWeek(d, time.Monday),
				"LAB",
			},
		)
		d = NextYear(d)
	}
	return holidays
}

func ColumbusDay(num int) []Holiday {
	// Second monday in October
	var holidays []Holiday
	d := PreviousMonthOfYear(PreviousYear(time.Now()), time.October)

	for len(holidays) < num {
		holidays = append(
			holidays,
			Holiday{
				"Columbus Day",
				NextWeek(NextDayOfWeek(d, time.Monday)),
				"COL",
			},
		)
		d = NextYear(d)
	}
	return holidays
}

func Thanksgiving(num int) []Holiday {
	// Fourth Thursday in November
	var holidays []Holiday
	d := PreviousMonthOfYear(PreviousYear(time.Now()), time.November)

	for len(holidays) < num {
		holidays = append(
			holidays,
			Holiday{
				"Thanksgiving",
				WeeksSince(NextDayOfWeek(d, time.Thursday), 3),
				"TGV",
			},
		)
		d = NextYear(d)
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
	// July 4th. Public holiday on Friday if it falls on a Saturday,
	// public holiday on the Monday if it falls on a Sunday
	var holidays []Holiday
	d := PreviousMonthOfYear(PreviousYear(time.Now()), time.July)
	for len(holidays) < num {
		july4th := d.AddDate(0, 0, 3)
		if july4th.Weekday() == time.Saturday {
			holidays = append(
				holidays,
				Holiday{
					"Independence Day observed",
					PreviousDay(july4th),
					"IND",
				},
			)
		} else if july4th.Weekday() == time.Sunday {
			holidays = append(
				holidays,
				Holiday{
					"Independence Day observed",
					NextDay(july4th),
					"IND",
				},
			)
		}
		holidays = append(
			holidays,
			Holiday{
				"Independence Day",
				july4th,
				"IND",
			},
		)
		d = NextYear(d)
	}
	return holidays
}

func VeteransDay(num int) []Holiday {
	// November 11th. Public holiday on Friday if it falls on a Saturday,
	// public holiday on the Monday if it falls on a Sunday
	var holidays []Holiday
	d := PreviousMonthOfYear(PreviousYear(time.Now()), time.November)
	for len(holidays) < num {
		november11th := d.AddDate(0, 0, 10)
		if november11th.Weekday() == time.Saturday {
			holidays = append(
				holidays,
				Holiday{
					"Veterans Day observed",
					PreviousDay(november11th),
					"VET",
				},
			)
		} else if november11th.Weekday() == time.Sunday {
			holidays = append(
				holidays,
				Holiday{
					"Veterans Day observed",
					NextDay(november11th),
					"VET",
				},
			)
		}
		holidays = append(
			holidays,
			Holiday{
				"Veterans Day",
				november11th,
				"VET",
			},
		)
		d = NextYear(d)
	}
	return holidays
}

func ChristmasDay(num int) []Holiday {
	// December 25th. Public holiday on Friday if it falls on a Saturday,
	// public holiday on the Monday if it falls on a Sunday
	var holidays []Holiday
	d := PreviousMonthOfYear(PreviousYear(time.Now()), time.December)
	for len(holidays) < num {
		december25th := d.AddDate(0, 0, 24)
		if december25th.Weekday() == time.Saturday {
			holidays = append(
				holidays,
				Holiday{
					"Christmas Day observed",
					PreviousDay(december25th),
					"CHR",
				},
			)
		} else if december25th.Weekday() == time.Sunday {
			holidays = append(
				holidays,
				Holiday{
					"Christmas Day observed",
					NextDay(december25th),
					"CHR",
				},
			)
		}
		holidays = append(
			holidays,
			Holiday{
				"Christmas Day",
				december25th,
				"CHR",
			},
		)
		d = NextYear(d)
	}
	return holidays
}
