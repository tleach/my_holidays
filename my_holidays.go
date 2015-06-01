package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	htmlTemplate "html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

func main() {
	m := martini.Classic()
	m.Get("/", serveHomepage)
	m.Get("/holidays.(?P<format>(json|ics))", serveHolidays)
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
	if error := tmpl.Execute(w, data); error != nil {
		fmt.Print(error)
	}
}

func renderJson(holidays []Holiday, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(holidays)
	w.Write(json)
}

func serveHomepage(w http.ResponseWriter, req *http.Request) {
	tmpl, error := htmlTemplate.ParseFiles("./templates/index.html")
	if error != nil {
		fmt.Print(error)
		return
	}
	data := struct {
		BaseUrl    string
		Generators []HolidayGenerator
	}{
		"https://" + req.Host,
		AllGenerators,
	}
	tmpl.Execute(w, data)
}

func serveHolidays(w http.ResponseWriter, req *http.Request, params martini.Params) {
	generators := AllGenerators
	if include := req.URL.Query().Get("include"); include != "" {
		generators = MatchCodes(generators, strings.Split(include, ","))
	}
	currentYear := time.Now().UTC().Year()
	holidays := GenerateAll(generators, currentYear-5, currentYear+25)

	if params["format"] == "ics" {
		renderCal(holidays, w)
	} else {
		renderJson(holidays, w)
	}
}
