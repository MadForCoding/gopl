package main

import (
	"ch6/geometry"
	"fmt"
)

func main(){

	// Test the distance between two points
	p := geometry.Point{1,1}
	q := geometry.Point{2,2}
	fmt.Println(p.Distance(q))

	// Test the set of points
	perim := geometry.Path{
		{1,1},
		{5,1},
		{5,4},
		{1,1},
	}

	fmt.Println(perim.Distance())
}
