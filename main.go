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
	// page := &Page{Title: "FormulaWhen"}
	// races := buildRaceSchedule()
	data := Data{
		Page:  Page{Title: "FormulaWhen"},
		Races: buildRaceSchedule(),
	}

	// fmt.Println(race.Name, race.Sessions[4].TimeToStart)
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

func (races RaceSchedule) getDurations() {
	for i, r := range races {
		for j, s := range r.Sessions {
			races[i].Sessions[j].TimeToStart = time.Until(s.StartTime)
		}
	}
}

func buildRaceSchedule() RaceSchedule {
	races := parseJson(readFile(raceDataPath))
	races.sortRaces()
	races.getDurations()
	return races
}

func (r RaceSchedule) sortRaces() {
	sort.Slice(r, func(i, j int) bool {
		// 4th session is race
		return r[i].Sessions[4].StartTime.Before(r[j].Sessions[4].StartTime)
	})
}
