// basename removes directory components and a .suffix.
// e.g. a => a, a.go => a, a/b/c.go => c
package main

import (
	"fmt"
	"os"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(basename(s))
	}
}

func basename(s string) string {
	// remove the last "/"
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// remove "."
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}