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
	Data  string
}

type Data struct {
	Page Page
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Page: Page{
			Title: "FormulaWhen",
			Data:  prepRaceData(raceDataPath)},
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

type Races []Race

type Race struct {
	Name     string    `json:"Name"`
	Sessions []Session `json:"Sessions"`
}

type Session struct {
	Name      string    `json:"Name"`
	StartTime time.Time `json:"StartTime"` // RFC3339
	EndTime   time.Time `json:"EndTime"`   // RFC3339
	NsToStart int
}

func prepRaceData(raceDataPath string) string {
	fileData := readFile(raceDataPath)
	races := buildRaces(fileData)
	return racesToJson(races)
}

func readFile(fp string) []byte {
	b, e := os.ReadFile(fp)
	if e != nil {
		panic(e)
	}
	return b
}

func parseJson(b []byte) Races {
	var races Races
	e := json.Unmarshal(b, &races)
	if e != nil {
		panic(e)
	}
	return races
}

// setTimeToSessions sets the time between time t and each session start time in
// races.
func (races Races) setTimeToSessions(t time.Time) {
	for i, r := range races {
		for j, s := range r.Sessions {
			races[i].Sessions[j].NsToStart = NsUntil(t, s.StartTime)
		}
	}
}

// secondsUntil returns the number of seconds between from until to if positive.
// If number of seconds is <= 0, returns  if positive. If number of seconds is
// <= 0, returns 0.
func NsUntil(from, to time.Time) int {
	s := int(to.Sub(from))
	if s < 0 {
		return 0
	}
	return s
}

func buildRaces(b []byte) Races {
	races := parseJson(b)
	races.sortRaces()
	races.setTimeToSessions(time.Now())
	races = removePastRaces(races, time.Now())
	return races
}

func racesToJson(r Races) string {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

func (r Races) sortRaces() {
	sort.Slice(r, func(i, j int) bool {
		return r[i].Sessions[0].StartTime.Before(r[j].Sessions[0].StartTime)
	})
}

// removePastRaces removes all races from a Races struct where the race
// start time is after t.
func removePastRaces(rr Races, t time.Time) Races {
	ret := Races{}
	for _, r := range rr {
		for _, s := range r.Sessions {
			if s.Name == "race" && s.EndTime.After(t) {
				ret = append(ret, r)
			}
		}
	}
	return ret
}
