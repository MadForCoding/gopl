package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}


func main() {
	var wheel = Wheel{
		Circle{
			Point{
				3,
				4,
			},
			5,
		},
		20,
	}

	fmt.Printf("%#v\n", wheel)

	wheel.X = 42
	fmt.Printf("%#v\n", wheel)
}