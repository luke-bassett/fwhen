package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type Page struct {
	Title string
}

type Data struct {
	Page  Page
	Races RaceSchedule
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Page:  Page{Title: "FormulaWhen"},
		Races: buildRaceSchedule(readFile(raceDataPath)),
	}

	t, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Starting on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

const raceDataPath = "data/2022_sessions.json"

type RaceSchedule []Race

type Race struct {
	Name     string    `json:"Name"`
	Sessions []Session `json:"Sessions"`
}

type Session struct {
	Name        string    `json:"Name"`
	StartTime   time.Time `json:"StartTime"` // RFC3339
	EndTime     time.Time `json:"EndTime"`   // RFC3339
	TimeToStart time.Duration
}

func readFile(fp string) []byte {
	b, e := os.ReadFile(fp)
	if e != nil {
		panic(e)
	}
	return b
}

func parseJson(b []byte) RaceSchedule {
	var races RaceSchedule
	e := json.Unmarshal(b, &races)
	if e != nil {
		panic(e)
	}
	return races
}

// Get duration between time t and each session start time in races.
func (races RaceSchedule) getDurations(t time.Time) {
	for i, r := range races {
		for j, s := range r.Sessions {
			races[i].Sessions[j].TimeToStart = s.StartTime.Sub(t)
		}
	}
}

func buildRaceSchedule(b []byte) RaceSchedule {
	races := parseJson(b)
	races.sortRaces()
	races.getDurations(time.Now())
	return races
}

func (r RaceSchedule) sortRaces() {
	sort.Slice(r, func(i, j int) bool {
		return r[i].Sessions[0].StartTime.Before(r[j].Sessions[0].StartTime)
	})
}
