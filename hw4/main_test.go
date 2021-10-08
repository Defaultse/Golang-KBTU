package main

import (
	"testing"
	"time"
)

func TestCountDays(t *testing.T) {
	type period struct {
		Begin time.Time
		End   time.Time
	}
	var tests = []struct {
		Input    period
		Expected float64
	}{
		{period{time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)}, 365},
		{period{time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC)}, 366},
	}

	for _, test := range tests {
		if output := countDays(test.Input.Begin, test.Input.End); output != test.Expected {
			t.Error("Test failes: got {}, expected {}, out {}", test.Input, test.Expected, output)
		}
	}
}
