package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageVariables struct {
	Date string
	Time string
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	// wow. Go uses layout numbers for time formatting. Seems like
	// a really stupid idea. I'm going to figure out how to fix this problem
	// for this, right now and then immediately forget how I did it because
	// it's stupid. I'm really curious what the reasoning behing this was.
	// Maybe one day I'll be curious enough to actually research it. Maybe
	// not.
	// Note to future me:
	// The official docs are not very helpful on this.
	// See this SO article:
	// https://stackoverflow.com/questions/25845172/parsing-date-string-in-golang
	// and this go package:
	// https://github.com/araddon/dateparse
	now := time.Now()
	indexVars := PageVariables{
		Date: now.Format("01-01-2020"),
		Time: now.Format("15:15:15"),
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, indexVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
