// oop project main.go
package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	//	var cp ColoredPoint
	//	cp.X = 1
	//	fmt.Println(cp.Point.X) // "1"
	//	cp.Point.Y = 2
	//	fmt.Println(cp.Y) // "2"
	//	red := color.RGBA{255, 0, 0, 255}
	//	blue := color.RGBA{0, 0, 255, 255}
	//	var p = ColoredPoint{Point{1, 1}, red}
	//	var q = ColoredPoint{Point{5, 4}, blue}
	//	fmt.Println(p.Distance(q.Point)) // "5"
	//	p.ScaleBy(2)
	//	q.ScaleBy(2)
	//	fmt.Println(p.Distance(q.Point))
	p := Point{1, 2}
	q := Point{4, 6}
	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
