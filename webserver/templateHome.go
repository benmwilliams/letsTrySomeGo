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
