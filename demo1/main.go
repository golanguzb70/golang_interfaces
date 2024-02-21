package main

import (
	"fmt"
	"math"
)

type shapes interface {
	Area()
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() {
	fmt.Println(c.Radius * c.Radius * math.Pi)
}

type Rectangle struct {
	Hight float64
	Width float64
}

func (r *Rectangle) Area() {
	fmt.Println(r.Hight * r.Width)
}

func PrintCircleArea(c *Circle) {
	c.Area()
}

func PrintRectangleArea(r *Rectangle) {
	r.Area()
}

func PrintArea(sh shapes) {
	sh.Area()
}

func main() {
	rectangle := Rectangle{
		Hight: 10,
		Width: 10,
	}

	circle := Circle{
		Radius: 10,
	}

	circle2 := &Circle{
		Radius: 12,
	}

	// PrintCircleArea(&circle)
	// PrintRectangleArea(&rectangle)

	PrintArea(&circle)
	PrintArea(&rectangle)

	shapeObjs := []shapes{&circle, &rectangle, circle2}
	for _, e := range shapeObjs {
		PrintArea(e)
	}
}
