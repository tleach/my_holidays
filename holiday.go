package main

import (
	"sort"
	"time"
)

type Holiday struct {
	Name         string
	Date         time.Time
	Code         string
	ObservedDate time.Time
}

func (h *Holiday) IsObservedOnDate() bool {
	return h.Date.Equal(h.ObservedDate)
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

func MatchCodes(generators []HolidayGenerator, codes []string) []HolidayGenerator {
	var matching []HolidayGenerator
	for _, g := range generators {
		for _, c := range codes {
			if c == g.Code {
				matching = append(matching, g)
			}
		}
	}
	return matching
}

type HolidayGenerator struct {
	DateForYear  func(year int) time.Time
	ObservedDate func(holidayDate time.Time) time.Time
	Name         string
	Description  string
	Code         string
}

func GenerateAll(generators []HolidayGenerator, fromYear, toYear int) []Holiday {
	var holidays []Holiday
	for _, g := range generators {
		holidays = append(holidays, g.Generate(fromYear, toYear)...)
	}
	sort.Sort(ByDate(holidays))
	return holidays
}

func (g *HolidayGenerator) Generate(fromYear, toYear int) []Holiday {
	var holidays []Holiday
	for i := fromYear; i <= toYear; i++ {
		d := g.DateForYear(i)
		h := Holiday{g.Name, d, g.Code, d}
		if g.ObservedDate != nil {
			h.ObservedDate = g.ObservedDate(d)
		}
		holidays = append(holidays, h)
	}
	return holidays
}
