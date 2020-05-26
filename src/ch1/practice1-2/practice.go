package main

import (
	"fmt"
	"os"
)

func main() {
	//var s, sep string

	for i, v := range os.Args[1:] {
		fmt.Println(i, ": ", v)
	}
}
