package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Get("/holidays", render)
	log.Printf("my_holidays listening on port 8020")
	err := http.ListenAndServe(":8027", m)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func render(w http.ResponseWriter, req *http.Request) {
	generators := []HolidayGenerator{
		NewYearsDay,
		MemorialDay,
	}

	var holidays []Holiday

	for _, generator := range generators {
		holidays = append(holidays, generator(10)...)
	}

	json, _ := json.Marshal(holidays)
	w.Write(json)
}
