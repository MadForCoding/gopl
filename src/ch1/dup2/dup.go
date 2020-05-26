// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from Stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	if length := len(os.Args[1:]); length == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range os.Args[1:] {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for word, number := range counts {
		if number > 1 {
			fmt.Printf("%d\t%s\n", number, word)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[strings.TrimSpace(input.Text())]++
	}
}
