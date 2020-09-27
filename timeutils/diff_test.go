package timeutils

import (
	"testing"
	"time"
)

func TestDiff(t *testing.T) {
	location, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("20060102", "20200521", location)
	t2, _ := time.ParseInLocation("20060102", "20200520", location)
	diff := Differ(t1, t2)
	if diff.Nanoseconds() != 24*time.Hour.Nanoseconds() {
		t.Error("differ wrong")
	} else {
		t.Log("differ success")
	}
}

func TestDifferWithDays(t *testing.T) {
	location, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("20060102", "20200521", location)
	t2, _ := time.ParseInLocation("20060102", "20200520", location)
	if DifferWithDays(t1, t2, 2) {
		t.Log("differ with days success")
	} else {
		t.Log("differ with days wrong")
	}
}
