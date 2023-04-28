package main

import (
	"testing"
	"time"
)

func TestInitCalendar(t *testing.T) {
	c, err := initCalendar()
	if err != nil {
		t.Error(err)
	}

	// Check that the sessions are in order
	lastSessionStartTime := time.Time{}
	for _, r := range c.Races {
		for _, s := range r.Sessions {
			if s.StartTime.Before(lastSessionStartTime) {
				t.Errorf("Session %s: %s starts before the previous session", r.Name, s.Name)
			}
		}
	}
}

// func TestFormat(t *testing.T) {
// 	c, err := initCalendar()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// Check that the formatted string is what we expect
// 	expected := `Formula 1 2023 -- All times UTC
// Page loaded: 2021-09-01 00:00:00
