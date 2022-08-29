package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "FormulaWhen"}
	t, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, p)
}

func main() {
	races := parseJson(readFile(raceDataPath))
	fmt.Println(races[0].Name)

	http.HandleFunc("/", handler)
	fmt.Println("Starting on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

const raceDataPath = "data/2022_sessions.json"

type RaceSchedule []struct {
	Name     string `json:"Name"`
	Sessions []struct {
		Name      string `json:"Name"`
		StartTime string `json:"StartTime"`
		EndTime   string `json:"EndTime"`
	} `json:"Sessions"`
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
