package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	want := "<!DOCTYPE html>"
	got := strings.Split(r.Body.String(), "\n")[0]
	if got != want {
		t.Errorf("handler returned unexpected first line of body: got %v want %v",
			got, want)
	}
}

func TestInitSessions(t *testing.T) {
	now := time.Now().UTC()
	calendar := &Calendar{
		Races: []Race{
			{
				Name:     "Race 1",
				Location: "Location 1",
				Sessions: []Session{
					{
						Name:      "Session 1",
						StartTime: now.Add(time.Hour),
					},
					{
						Name:      "Session 2",
						StartTime: now.Add(2 * time.Hour),
					},
				},
			},
		},
		ReferenceTime: now,
	}
	calendar.initSessions()
	s1Got := calendar.Races[0].Sessions[0].TimeUntil
	if s1Got != time.Hour {
		t.Errorf("Session 1 time until is incorrect: got %v want %v",
			s1Got, time.Hour)
	}
	s2Got := calendar.Races[0].Sessions[0].TimeUntil
	if s2Got != time.Hour {
		t.Errorf("Session 2 time until is incorrect: got %v want %v",
			s2Got, time.Hour)
	}
}

func TestFormatDuration(t *testing.T) {
	// Duration is positive
	duration1 := time.Hour*24 + time.Hour*2 + time.Minute*30
	expectedResult1 := "  1d 02h 30m 00s"
	result1 := formatDuration(duration1)
	if result1 != expectedResult1 {
		t.Errorf("Expected '%s', got '%s'", expectedResult1, result1)
	}

	// Duration is zero
	duration2 := time.Duration(0)
	expectedResult2 := "  0d 00h 00m 00s"
	result2 := formatDuration(duration2)
	if result2 != expectedResult2 {
		t.Errorf("Expected '%s', got '%s'", expectedResult2, result2)
	}

	// Duration is negative
	duration3 := -time.Hour*3 - time.Minute*15
	expectedResult3 := "     ---"
	result3 := formatDuration(duration3)
	if result3 != expectedResult3 {
		t.Errorf("Expected '%s', got '%s'", expectedResult3, result3)
	}
}
