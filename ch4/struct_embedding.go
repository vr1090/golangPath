package main

import "fmt"

type Point struct {
	X, Y int
}

type Wheels struct {
	Point
	Radius int
}

func main() {
	var w Wheels
	w.X = 100
	w.Y = 90
	w.Radius = 88

	fmt.Println("hahha ", w)

	var k = Wheels{
		Point:  Point{4, 5},
		Radius: 90,
	}

	fmt.Println(k)
}
