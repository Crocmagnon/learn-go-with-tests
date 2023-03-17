package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), 3 * math.Pi / 2},
		{simpleTime(0, 0, 15), math.Pi / 2},
		{simpleTime(0, 0, 7), 7 * math.Pi / 30},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondsInRadians(test.time)
			if !roughlyEqualFloat64(got, test.angle) {
				t.Fatalf("Wanted %v radians, got %v", test.angle, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * math.Pi / (30 * 60)},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 45, 0), 3 * math.Pi / 2},
		{simpleTime(0, 15, 0), math.Pi / 2},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := minutesInRadians(test.time)
			if !roughlyEqualFloat64(got, test.angle) {
				t.Fatalf("Wanted %v radians, got %v", test.angle, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), 3 * math.Pi / 2},
		{simpleTime(0, 1, 30), math.Pi / (6 * 60 * 60 / 90)},
	}

	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := hoursInRadians(test.time)
			if !roughlyEqualFloat64(got, test.angle) {
				t.Fatalf("Wanted %v radians, got %v", test.angle, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondHandPoint(test.time)
			if !roughlyEqualPoint(got, test.point) {
				t.Fatalf("Wanted %v Point, got %v", test.point, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}
	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := minuteHandPoint(test.time)
			if !roughlyEqualPoint(got, test.point) {
				t.Fatalf("Wanted %v Point, got %v", test.point, got)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}
	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := hourHandPoint(test.time)
			if !roughlyEqualPoint(got, test.point) {
				t.Fatalf("Wanted %v Point, got %v", test.point, got)
			}
		})
	}
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 24, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
