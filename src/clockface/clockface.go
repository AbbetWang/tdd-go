package clockface

import (
	"math"
	"time"
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

const secondsInHalfClock = 30
const secondInClock = 2 * secondsInHalfClock
const minutesInHalfClock = 30
const minutesInClock = 2 * minutesInHalfClock
const hoursInHalfClock = 6
const hoursInClock = 2 * hoursInHalfClock

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}
func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians((t)))
}
func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}
func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}
