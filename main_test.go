package main

import (
	"testing"
	"time"
)

const sample_json = `
[{
    "Name": "brazil",
   		"Sessions": [
      		{
				"Name": "fp1",
				"StartTime": "2022-11-11T15:30:00Z",
				"EndTime": "2022-11-11T16:30:00Z"
      }
	]
},
{
    "Name": "mexico",
   		"Sessions": [
	        {
				"Name": "fp2",
				"StartTime": "2022-10-28T18:00:00Z",
				"EndTime": "2022-10-28T19:00:00Z"
			}
	]
}]`

func TestParseJson(t *testing.T) {
	j := parseJson([]byte(sample_json))
	got := j[1].Name
	want := "mexico"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
	got = j[0].Sessions[0].Name
	want = "fp1"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSortRaces(t *testing.T) {
	races := parseJson([]byte(sample_json))

	races.sortRaces()
	got := races[0].Name
	want := "mexico"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSetTimeToSessions(t *testing.T) {
	races := parseJson([]byte(sample_json))
	// 24 hours
	tt := time.Date(2022, 10, 27, 18, 0, 0, 0, time.UTC)
	races.setTimeToSessions(tt)

	got := races[1].Sessions[0].NsToStart
	want := int(time.Duration(24) * time.Hour)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
	// already happened
	tt = time.Date(2050, 6, 14, 12, 23, 44, 0, time.UTC)
	races.setTimeToSessions(tt)
	got = races[1].Sessions[0].NsToStart
	want = 0
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestBuildSchedule(t *testing.T) {
	races := buildRaceSchedule([]byte(sample_json))
	got := races[1].Sessions[0].Name
	want := "fp1"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
	got2 := races[0].Sessions[0].EndTime
	want2, _ := time.Parse(time.RFC3339, "2022-10-28T19:00:00Z")
	if got2 != want2 {
		t.Errorf("got %s, wanted %s", got2, want2)
	}
}
