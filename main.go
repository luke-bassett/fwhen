package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const scheduleJsonFile = "2023-f1-schedule.json"

type Calendar struct {
	Races         []Race
	ReferenceTime time.Time
}

type Race struct {
	Name     string `json:"gp"`
	Location string `json:"location"`
	Sessions []Session
}

type Session struct {
	Name            string    `json:"session"`
	Cancelled       bool      `json:"cancelled"`
	StartTime       time.Time `json:"start"`
	TimeUntil       time.Duration
	TimeUntilString string
}

func initCalendar() (*Calendar, error) {
	cal := Calendar{ReferenceTime: time.Now().UTC()}
	rawJson, err := os.ReadFile(scheduleJsonFile)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(rawJson, &cal.Races); err != nil {
		return nil, err
	}
	cal.initSessions()
	return &cal, nil
}

func (cal *Calendar) initSessions() {
	for r := range cal.Races {
		for s := range cal.Races[r].Sessions {
			session := &cal.Races[r].Sessions[s]
			session.TimeUntil = session.StartTime.Sub(cal.ReferenceTime)
			if session.TimeUntil < 0 {
				session.TimeUntilString = "     ---"
			} else if session.Cancelled {
				session.TimeUntilString = "Cancelled"
			} else {
				session.TimeUntilString = formatDuration(session.TimeUntil)
			}
		}
	}
}

func formatDuration(duration time.Duration) string {
	if duration < 0 {
		return "     ---"
	}
	return fmt.Sprintf("%3dd %02dh %02dm %02ds",
		int(duration.Hours()/24),
		int(duration.Hours())%24,
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}

func handler(w http.ResponseWriter, r *http.Request) {
	cal, err := initCalendar()
	if err != nil {
		log.Fatal(err)
	}
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, cal)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.Handle("/js/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
