// Dup1 prints the text of each line that appear more than
// once in the standard input, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// control + D to stop scan
	for input.Scan() {
		counts[input.Text()]++
	}

	for word, number := range counts {
		if number > 1 {
			fmt.Printf("%d\t%s\n", number, word)
		}
	}

}
