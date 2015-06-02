package main

import (
	"time"
)

func USObservationDate(date time.Time) time.Time {
	if date.Weekday() == time.Saturday {
		return PreviousDay(date)
	} else if date.Weekday() == time.Sunday {
		return NextDay(date)
	}
	return date
}

var NewYearsDay HolidayGenerator = HolidayGenerator{
	Name:         "New Year's Day",
	Description:  "January 1st",
	Code:         "NEW",
	ObservedDate: USObservationDate,
	DateForYear: func(year int) time.Time {
		return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	},
}

var MartinLutherKingJuniorDay HolidayGenerator = HolidayGenerator{
	Name:        "Martin Luther King Jr. Day",
	Description: "Third Monday in January",
	Code:        "MLK",
	DateForYear: func(year int) time.Time {
		return WeeksSince(NextDayOfWeek(time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC), time.Monday), 2)
	},
}

var PresidentsDay HolidayGenerator = HolidayGenerator{
	Name:        "Presidents Day",
	Description: "Third Monday of February",
	Code:        "PRS",
	DateForYear: func(year int) time.Time {
		d := time.Date(year, time.February, 1, 0, 0, 0, 0, time.UTC)
		return WeeksSince(NextDayOfWeek(d, time.Monday), 2)
	},
}

var MemorialDay HolidayGenerator = HolidayGenerator{
	Name:        "Memorial Day",
	Description: "Last Monday in May",
	Code:        "MEM",
	DateForYear: func(year int) time.Time {
		return PreviousDayOfWeek(
			EndOfMonth(
				time.Date(year, time.May, 1, 0, 0, 0, 0, time.UTC),
			),
			time.Monday)
	},
}

var IndependenceDay HolidayGenerator = HolidayGenerator{
	Name:         "Independence Day",
	Description:  "July 4th",
	Code:         "IND",
	ObservedDate: USObservationDate,
	DateForYear: func(year int) time.Time {
		return time.Date(year, time.July, 4, 0, 0, 0, 0, time.UTC)
	},
}

var LaborDay HolidayGenerator = HolidayGenerator{
	Name:        "Labor Day",
	Description: "First Monday in September",
	Code:        "LAB",
	DateForYear: func(year int) time.Time {
		return NextDayOfWeek(
			time.Date(year, time.September, 1, 0, 0, 0, 0, time.UTC), time.Monday)
	},
}

var ColumbusDay HolidayGenerator = HolidayGenerator{
	Name:        "Columbus Day",
	Description: "Second Monday in October",
	Code:        "COL",
	DateForYear: func(year int) time.Time {
		return NextWeek(NextDayOfWeek(time.Date(year, time.October, 1, 0, 0, 0, 0, time.UTC), time.Monday))
	},
}

var Thanksgiving HolidayGenerator = HolidayGenerator{
	Name:        "Thanksgiving",
	Description: "Fourth Thursday in November",
	Code:        "TGV",
	DateForYear: func(year int) time.Time {
		return WeeksSince(NextDayOfWeek(time.Date(year, time.November, 1, 0, 0, 0, 0, time.UTC), time.Thursday), 3)
	},
}

var VeteransDay HolidayGenerator = HolidayGenerator{
	Name:         "Veterans Day",
	Description:  "November 11th",
	Code:         "VET",
	ObservedDate: USObservationDate,
	DateForYear: func(year int) time.Time {
		return time.Date(year, time.November, 11, 0, 0, 0, 0, time.UTC)
	},
}

var ChristmasDay HolidayGenerator = HolidayGenerator{
	Name:         "Christmas Day",
	Description:  "December 25th",
	Code:         "XMS",
	ObservedDate: USObservationDate,
	DateForYear: func(year int) time.Time {
		return time.Date(year, time.December, 25, 0, 0, 0, 0, time.UTC)
	},
}

var AllGenerators []HolidayGenerator = []HolidayGenerator{
	NewYearsDay, MartinLutherKingJuniorDay, PresidentsDay,
	MemorialDay, IndependenceDay, LaborDay, ColumbusDay,
	Thanksgiving, VeteransDay, ChristmasDay,
}
