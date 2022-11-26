package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
    X float64
    Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}  // scale
	p = Point{p.X, -p.Y}  // flip
	p = Point{p.X + 150, p.Y + 150}  // translate
	
	return p

    return Point{150, 60}
}


func secondsInRadians(t time.Time) float64 {
	// by structuring the equation so we don't divide down and then multiply up with Pi,
	// we are able to preserve the float64 equality, which we otherwise wouldve lost. 
	// aka Pi / 30 != 30 * (Pi / 30)
	return (math.Pi / (30 / float64(t.Second())))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}