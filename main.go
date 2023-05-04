package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

const scheduleJsonFile = "static/2023-f1-schedule.json"
const dateFormat = "2006-01-02 15:04:05"

type Calendar struct {
	Races         []Race
	ReferenceTime time.Time
	Text          string
}

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

func initCalendar() (*Calendar, error) {
	c := Calendar{ReferenceTime: time.Now().UTC()}
	rawJson, err := os.ReadFile(scheduleJsonFile)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(rawJson, &c.Races); err != nil {
		return nil, err
	}
	c.Text = c.format()
	return &c, nil
}

func (c Calendar) format() string {
	var str string
	for _, r := range c.Races {
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
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, c)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
