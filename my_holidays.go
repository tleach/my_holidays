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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8027"
	}
	m := martini.Classic()
	m.Get("/", serveHomepage)
	m.Get("/holidays.(?P<format>(json|ics))", serveHolidays)
	log.Printf("my_holidays listening on port %v", port)
	err := http.ListenAndServe(":"+port, m)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var _tmpl *template.Template

func calTemplate() (*template.Template, error) {
	cwd, _ := os.Getwd()
	if _tmpl == nil {
		fmt.Print("LOad")
		var err error
		_tmpl, err = template.ParseFiles(filepath.Join(cwd, "./templates/holidays.ics"))
		if err != nil {
			return nil, err
		}
	}
	return _tmpl, nil
}

func renderCal(title string, holidays []Holiday, w http.ResponseWriter) error {
	tmpl, err := calTemplate()
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/calendar")
	w.Header().Set("Content-Disposition", "attachment; filename=\"holidays.ics\"")
	data := struct {
		Title    string
		Holidays []Holiday
	}{
		title,
		holidays,
	}
	if err := tmpl.Execute(w, data); err != nil {
		return err
	}
	return nil
}

func renderJson(title string, holidays []Holiday, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(holidays)
	if err != nil {
		return err
	}
	w.Write(json)
	return nil
}

func serveHomepage(w http.ResponseWriter, req *http.Request) {
	tmpl, error := htmlTemplate.ParseFiles("./templates/index.html")
	if error != nil {
		fmt.Print(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := struct {
		BaseUrl           string
		USFederalHolidays []HolidayGenerator
		USOtherHolidays   []HolidayGenerator
	}{
		"http://" + req.Host,
		USFederalHolidays,
		USOtherHolidays,
	}
	tmpl.Execute(w, data)
}

func serveHolidays(w http.ResponseWriter, req *http.Request, params martini.Params) {
	generators := AllGenerators
	if include := req.URL.Query().Get("include"); include != "" {
		generators = MatchCodes(generators, strings.Split(include, ","))
	}
	title := req.URL.Query().Get("title")
	if title == "" {
		title = "My Holidays"
	}
	currentYear := time.Now().UTC().Year()
	holidays := GenerateAll(generators, currentYear-5, currentYear+25)

	var err error
	if params["format"] == "ics" {
		err = renderCal(title, holidays, w)
	} else {
		err = renderJson(title, holidays, w)
	}
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
