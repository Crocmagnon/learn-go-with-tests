package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// A point represents a two-dimensional cartesian coordinate
type point struct {
	X, Y float64
}

func secondsInRadians(t time.Time) float64 {
	seconds := float64(t.Second())
	return math.Pi / (secondsInHalfClock / seconds)
}

func minutesInRadians(t time.Time) float64 {
	minutes := float64(t.Minute())
	return (secondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / minutes))
}

func hoursInRadians(t time.Time) float64 {
	hours := float64(t.Hour() % hoursInClock)
	return (minutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / hours))
}

func angleToPoint(angle float64) point {
	return point{math.Sin(angle), math.Cos(angle)}
}

func secondHandPoint(t time.Time) point {
	angle := secondsInRadians(t)
	return angleToPoint(angle)
}

func minuteHandPoint(t time.Time) point {
	angle := minutesInRadians(t)
	return angleToPoint(angle)
}

func hourHandPoint(t time.Time) point {
	angle := hoursInRadians(t)
	return angleToPoint(angle)
}
