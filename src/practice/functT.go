package main

import (
	"fmt"
	"os"
)

func main(){
	linenum, name := 12, "undefined"
	errorF(linenum,"undefined: %s", name)
}

func errorF(numLine int, stringFormat string, args ...interface{}){
	fmt.Fprintf(os.Stderr, "Line %d: ", numLine)
	fmt.Fprintf(os.Stderr, stringFormat, args)
	fmt.Fprintln(os.Stderr)
}