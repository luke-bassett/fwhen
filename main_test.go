package main

import "testing"

const sample_json = `
[
  {
    "Name": "brazil",
    "Sessions": [
      {
        "Name": "fp1",
        "StartTime": "2022-11-11T15:30:00Z",
        "EndTime": "2022-11-11T16:30:00Z"
      }
	]
  }
]`

func TestParseJson(t *testing.T) {
	got := parseJson([]byte(sample_json))[0].Name
	want := "brazil"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
