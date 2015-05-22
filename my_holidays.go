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

	json, _ := json.Marshal(MemorialDay(10))
	w.Write(json)
}
