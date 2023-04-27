package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const scheduleJsonFile = "2023-f1-schedule.json"
const dateFormat = "2006-01-02 15:04:05"

type Calendar []Race

type Race struct {
	Name     string    `json:"gp"`
	Location string    `json:"location"`
	Sessions []Session `json:"sessions"`
}

type Session struct {
	Name      string    `json:"session"`
	StartTime time.Time `json:"start"`
	EndTime   time.Time `json:"end"`
}

func initCalendar() (Calendar, error) {
	var c Calendar
	rawJson, err := os.ReadFile(scheduleJsonFile)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(rawJson, &c); err != nil {
		return nil, err
	}
	return c, nil
}

func (c Calendar) format() string {
	var str string
	for _, r := range c {
		str += fmt.Sprintf("%v - %v\n", r.Name, r.Location)
		for _, s := range r.Sessions {
			str += fmt.Sprintf("  %-12v%v\n", s.Name, s.StartTime.Format(dateFormat))
		}
	}
	return str
}

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := initCalendar()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Formula 1 2023 -- All times UTC\n")
	fmt.Fprintf(w, "Page loaded: %+v\n\n", time.Now().UTC().Format(dateFormat))
	fmt.Fprint(w, c.format())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
