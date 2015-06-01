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

func NewYearsDay() HolidayGenerator {
	return HolidayGenerator{
		Name:         "New Year's Day",
		Description:  "First day of the year",
		Code:         "NEW",
		ObservedDate: USObservationDate,
		DateForYear: func(year int) time.Time {
			return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		},
	}
}

func MemorialDay() HolidayGenerator {
	// Memorial Day falls on the last monday in May
	return HolidayGenerator{
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
}

func LaborDay() HolidayGenerator {
	return HolidayGenerator{
		Name:        "Labor Day",
		Description: "First Monday in September",
		Code:        "LAB",
		DateForYear: func(year int) time.Time {
			return NextDayOfWeek(
				time.Date(year, time.September, 1, 0, 0, 0, 0, time.UTC), time.Monday)
		},
	}
}

func ColumbusDay() HolidayGenerator {
	return HolidayGenerator{
		Name:        "Columbus Day",
		Description: "Second Monday in October",
		Code:        "COL",
		DateForYear: func(year int) time.Time {
			return NextWeek(NextDayOfWeek(time.Date(year, time.October, 1, 0, 0, 0, 0, time.UTC), time.Monday))
		},
	}
}

func Thanksgiving() HolidayGenerator {
	return HolidayGenerator{
		Name:        "Thanksgiving",
		Description: "Fourth Thursday in November",
		Code:        "TGV",
		DateForYear: func(year int) time.Time {
			return WeeksSince(NextDayOfWeek(time.Date(year, time.November, 1, 0, 0, 0, 0, time.UTC), time.Thursday), 3)
		},
	}
}

func MartinLutherKingJuniorDay() HolidayGenerator {
	return HolidayGenerator{
		Name:        "Martin Luther King Jr. Day",
		Description: "Third Monday in January",
		Code:        "MLK",
		DateForYear: func(year int) time.Time {
			return WeeksSince(NextDayOfWeek(time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC), time.Monday), 2)
		},
	}
}

func PresidentsDay() HolidayGenerator {
	return HolidayGenerator{
		Name:        "Presidents Day",
		Description: "Third Monday of February",
		Code:        "PRS",
		DateForYear: func(year int) time.Time {
			d := time.Date(year, time.February, 1, 0, 0, 0, 0, time.UTC)
			return WeeksSince(NextDayOfWeek(d, time.Monday), 2)
		},
	}
}

func IndependenceDay() HolidayGenerator {
	return HolidayGenerator{
		Name:         "Independence Day",
		Description:  "The Fourth of July",
		Code:         "IND",
		ObservedDate: USObservationDate,
		DateForYear: func(year int) time.Time {
			return time.Date(year, time.July, 4, 0, 0, 0, 0, time.UTC)
		},
	}
}

func VeteransDay() HolidayGenerator {
	return HolidayGenerator{
		Name:         "Veterans Day",
		Description:  "The Eleventh of November",
		Code:         "VET",
		ObservedDate: USObservationDate,
		DateForYear: func(year int) time.Time {
			return time.Date(year, time.November, 11, 0, 0, 0, 0, time.UTC)
		},
	}
}

func ChristmasDay() HolidayGenerator {
	return HolidayGenerator{
		Name:        "Christmas Day",
		Description: "The Twenty-Fifth of December",
		Code:        "XMS",
		DateForYear: func(year int) time.Time {
			return time.Date(year, time.December, 25, 0, 0, 0, 0, time.UTC)
		},
	}
}

func AllGenerators() []HolidayGenerator {
	return []HolidayGenerator{
		NewYearsDay(), MemorialDay(), LaborDay(), ColumbusDay(),
		Thanksgiving(), MartinLutherKingJuniorDay(), PresidentsDay(),
		IndependenceDay(), VeteransDay(), ChristmasDay(),
	}
}
