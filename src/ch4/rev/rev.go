package main

import "fmt"

func main() {
	s := []int{1,2,3,4,5,6,7,8,9}
	reverse(s)
	fmt.Println(s)
}

// reverse reverses a slice of ints in place
func reverse(s []int) {
	for left, right := 0, len(s) -1; left < right; left, right = left+1, right - 1{
		s[left], s[right] = s[right], s[left]
	}
}
