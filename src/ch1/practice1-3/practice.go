package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s1, sep string

	// test the duration without using strings.Join
	start := time.Now()
	for _, v := range os.Args[1:] {
		s1 += sep + v
		sep = " "
	}
	end := time.Now()
	fmt.Println("time duration without using strings.Join: ", end.Sub(start))


	// test while using strings.Join
	start = time.Now()
	strings.Join(os.Args[1:], " ")
	end = time.Now()
	fmt.Println("time duration using strings.Join: ", end.Sub(start))
}
