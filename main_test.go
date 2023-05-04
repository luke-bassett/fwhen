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

func TestFormat(t *testing.T) {
	c, err := initCalendar()
	if err != nil {
		t.Error(err)
	}
	formatted := c.format()
	nLines := len(strings.Split(strings.TrimSpace(formatted), "\n"))
	expectedLines := len(c.Races) * 6
	if nLines != expectedLines {
		t.Errorf("Have %v lines, but expect %v", nLines, expectedLines)
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
