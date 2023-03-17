package clockface

import (
	"math"
	"time"
)

// A Point represents a two-dimensional cartesian coordinate
type Point struct {
	X, Y float64
}

func secondsInRadians(t time.Time) float64 {
	seconds := float64(t.Second())
	return math.Pi / (30 / seconds)
}

func minutesInRadians(t time.Time) float64 {
	minutes := float64(t.Minute())
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / minutes))
}

func angleToPoint(angle float64) Point {
	return Point{math.Sin(angle), math.Cos(angle)}
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	return angleToPoint(angle)
}

func minuteHandPoint(t time.Time) Point {
	angle := minutesInRadians(t)
	return angleToPoint(angle)
}
