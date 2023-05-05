package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const scheduleJsonFile = "static/2023-f1-schedule.json"

type Calendar struct {
	Races         []Race
	ReferenceTime time.Time
}

type Race struct {
	Name     string    `json:"gp"`
	Location string    `json:"location"`
	Sessions []Session `json:"sessions"`
}

type Session struct {
	Name      string    `json:"session"`
	StartTime time.Time `json:"start"`
	TimeUntil time.Duration
}

func initCalendar() (*Calendar, error) {
	c := Calendar{ReferenceTime: time.Now().UTC()}
	rawJson, err := os.ReadFile(scheduleJsonFile)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(rawJson, &c.Races); err != nil {
		return nil, err
	}
	c.CalculateTimeUntil()
	return &c, nil
}

func (c *Calendar) CalculateTimeUntil() {
	for _, r := range c.Races {
		for i := range r.Sessions {
			s := &r.Sessions[i]
			s.TimeUntil = s.StartTime.Sub(c.ReferenceTime)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := initCalendar()
	if err != nil {
		log.Fatal(err)
	}
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, c)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
