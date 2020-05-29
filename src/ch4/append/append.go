package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(s []int, v int) []int {
	var z []int
	zLen := len(s) + 1
	if zLen <= cap(s) {
		z = s[:zLen]
	} else {
		zCap := zLen
		if zCap < 2 * len(s) {
			zCap = 2 * len(s)
		}
		z = make([]int, zLen, zCap)
		copy(z, s)
	}
	z[len(s)] = v
	return z
}


