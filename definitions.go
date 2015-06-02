package main

import (
	"time"
)

// Calculates and returns the date on which the given holiday
// date is actually observed. In the US, if a designated
// holiday falls on a Saturday, then the preceding Friday is
// taken as the "observed" holiday - a non-working day.
// If the designated holiday falls on a Sunday, it is
// observed on the following Monday.
func USObservationDate(date time.Time) time.Time {
	if date.Weekday() == time.Saturday {
		return PreviousDay(date)
	} else if date.Weekday() == time.Sunday {
		return NextDay(date)
	}
	return date
}

// US FEDERAL HOLIDAYS

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

// US Other Holidays

var GroundhogDay HolidayGenerator = HolidayGenerator{
	Name:        "Groundhog Day",
	Description: "February 2nd",
	Code:        "GHG",
	DateForYear: func(year int) time.Time {
		return time.Date(year, 2, 2, 0, 0, 0, 0, time.UTC)
	},
}

var ValentinesDay HolidayGenerator = HolidayGenerator{
	Name:        "Valentine's Day",
	Description: "February 14th",
	Code:        "VAL",
	DateForYear: func(year int) time.Time {
		return time.Date(year, 2, 14, 0, 0, 0, 0, time.UTC)
	},
}

var EarthDay HolidayGenerator = HolidayGenerator{
	Name:        "Earth Day",
	Description: "April 22nd",
	Code:        "EAR",
	DateForYear: func(year int) time.Time {
		return time.Date(year, 4, 22, 0, 0, 0, 0, time.UTC)
	},
}

var ArborDay HolidayGenerator = HolidayGenerator{
	Name:        "Arbor Day",
	Description: "Last Friday in April",
	Code:        "ARB",
	DateForYear: func(year int) time.Time {
		return PreviousDayOfWeek(
			EndOfMonth(
				time.Date(year, 4, 1, 0, 0, 0, 0, time.UTC)),
			time.Friday)
	},
}

var MothersDay HolidayGenerator = HolidayGenerator{
	Name:        "Mothers Day",
	Description: "Second Sunday in May",
	Code:        "MOM",
	DateForYear: func(year int) time.Time {
		return NextWeek(
			NextDayOfWeek(
				time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC), time.Sunday))
	},
}

var FlagDay HolidayGenerator = HolidayGenerator{
	Name:        "Flag Day",
	Description: "June 14th",
	Code:        "FLG",
	DateForYear: func(year int) time.Time {
		return time.Date(year, 6, 14, 0, 0, 0, 0, time.UTC)
	},
}

var FathersDay HolidayGenerator = HolidayGenerator{
	Name:        "Fathers Day",
	Description: "Third Sunday in June",
	Code:        "POP",
	DateForYear: func(year int) time.Time {
		return WeeksSince(
			NextDayOfWeek(
				time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC),
				time.Sunday),
			2)
	},
}

var PatriotDay HolidayGenerator = HolidayGenerator{
	Name:        "Patriot Day",
	Description: "September 11th",
	Code:        "PAT",
	DateForYear: func(year int) time.Time {
		return time.Date(year, 9, 11, 0, 0, 0, 0, time.UTC)
	},
}

var Halloween HolidayGenerator = HolidayGenerator{
	Name:        "Halloween",
	Description: "October 31st",
	Code:        "HAL",
	DateForYear: func(year int) time.Time {
		return time.Date(year, 10, 31, 0, 0, 0, 0, time.UTC)
	},
}

var PearlHarborDay HolidayGenerator = HolidayGenerator{
	Name:        "Pearl Harbor Day",
	Description: "December 7th",
	Code:        "PHB",
	DateForYear: func(year int) time.Time {
		return time.Date(year, 12, 7, 0, 0, 0, 0, time.UTC)
	},
}

var USFederalHolidays []HolidayGenerator = []HolidayGenerator{
	NewYearsDay, MartinLutherKingJuniorDay, PresidentsDay,
	MemorialDay, IndependenceDay, LaborDay, ColumbusDay,
	Thanksgiving, VeteransDay, ChristmasDay,
}

var USOtherHolidays []HolidayGenerator = []HolidayGenerator{
	GroundhogDay, ValentinesDay, EarthDay, ArborDay, MothersDay,
	FlagDay, FathersDay, PatriotDay, Halloween, PearlHarborDay,
}

var AllGenerators []HolidayGenerator = append(USFederalHolidays, USOtherHolidays...)
