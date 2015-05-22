package main

import (
	"time"
)

func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	return BeginningOfMonth(t).AddDate(0, 1, 0).AddDate(0, 0, -1)
}

func BeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func NextDayOfWeek(t time.Time, w time.Weekday) time.Time {
	diff := w - t.Weekday()
	if diff < 0 {
		diff = diff + 7
	}
	return BeginningOfDay(t.AddDate(0, 0, int(diff)))
}

func PreviousDayOfWeek(t time.Time, w time.Weekday) time.Time {
	diff := w - t.Weekday()
	if diff > 0 {
		diff = diff - 7
	}
	return BeginningOfDay(t.AddDate(0, 0, int(diff)))
}

func NextMonth(t time.Time) time.Time {
	return t.AddDate(0, 1, 0)
}

func PreviousMonth(t time.Time) time.Time {
	return t.AddDate(0, -1, 0)
}

func NextYear(t time.Time) time.Time {
	return t.AddDate(1, 0, 0)
}

func PreviousYear(t time.Time) time.Time {
	return t.AddDate(-1, 0, 0)
}
