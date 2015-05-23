package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

func main() {
	m := martini.Classic()
	m.Get("/holidays.(?P<format>(json|ics))", render)
	log.Printf("my_holidays listening on port 8027")
	err := http.ListenAndServe(":8027", m)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func renderCal(holidays []Holiday, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/calendar")
	w.Header().Set("Content-Disposition", "attachment; filename=\"holidays.ics\"")
	cwd, _ := os.Getwd()
	tmpl, error := template.ParseFiles(filepath.Join(cwd, "./templates/holidays.ics"))
	if error != nil {
		fmt.Print(error)
		return
	}
	data := struct {
		Holidays []Holiday
	}{
		holidays,
	}
	tmpl.Execute(w, data)
}

func renderJson(holidays []Holiday, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(holidays)
	w.Write(json)
}

func render(w http.ResponseWriter, req *http.Request, params martini.Params) {
	availableGenerators := map[string]HolidayGenerator{
		"NEW": NewYearsDay,
		"MEM": MemorialDay,
		"LAB": LaborDay,
		"COL": ColumbusDay,
		"TGV": Thanksgiving,
		"MLK": MartinLutherKingJuniorDay,
		"PRS": PresidentsDay,
		"IND": IndependenceDay,
		"VET": VeteransDay,
		"CHR": ChristmasDay,
	}

	var generators []HolidayGenerator
	if include := req.URL.Query().Get("include"); include != "" {
		for _, code := range strings.Split(include, ",") {
			generators = append(generators, availableGenerators[code])
		}
	} else {
		for _, v := range availableGenerators {
			generators = append(generators, v)
		}
	}

	var holidays []Holiday

	for _, generator := range generators {
		holidays = append(holidays, generator(10)...)
	}
	sort.Sort(ByDate(holidays))

	if params["format"] == "ics" {
		renderCal(holidays, w)
	} else {
		renderJson(holidays, w)
	}
}
