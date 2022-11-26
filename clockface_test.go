package clockface

import (
	"math"
	"testing"
	"time"
)

func simpleTime (hour, min, sec int) time.Time {
	return time.Date(2021, time.November, 26, hour, min, sec, 0, time.UTC)
}

func testName(time time.Time) string {
	return time.Format("15:04:05")
}


func TestSecondsInRadians(t *testing.T) {

	cases := []struct {
		time time.Time
		angle float64
	} {
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},

	}
	
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)

			if got != c.angle {
				t.Fatalf("got %v want%v", got, c.angle)
			}
		})
	}

}