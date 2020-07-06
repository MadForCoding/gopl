package geometry

import "math"

type Point struct{
	X, Y float64
}

// Traditional function to calculate the distance between two points
func Distance(p, q Point) float64{
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}

// same thing, but as method of the Point type
func (p Point) Distance(q Point) float64{
	return math.Hypot(p.X - q.X, p.Y - q.Y)
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

func (path Path) Distance() float64{
	sum := 0.0
	for index := range path{
		if index > 0 {
			sum += path[index-1].Distance(path[index])
		}
	}

	return sum
}



