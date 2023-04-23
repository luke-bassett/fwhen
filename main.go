package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const scheduleJsonFile = "2023-f1-schedule.json"

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

func loadCalendar() (Calendar, error) {
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

func main() {
	c, err := loadCalendar()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", c)
}
