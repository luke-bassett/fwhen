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

const scheduleJsonFile = "static/2023-f1-schedule.json"

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
	StartTime       time.Time `json:"start"`
	TimeUntil       time.Duration
	TimeUntilString string
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
	for r := range c.Races {
		for s := range c.Races[r].Sessions {
			c.Races[r].Sessions[s].TimeUntil = c.Races[r].Sessions[s].StartTime.Sub(c.ReferenceTime)
			c.Races[r].Sessions[s].TimeUntilString = formatDuration(c.Races[r].Sessions[s].TimeUntil)
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
