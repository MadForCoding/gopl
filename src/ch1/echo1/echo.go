// Echo1 prints its command-line arguements
package main

import (
	"fmt"
	"os"
)

func main() {
	// If the variable not initalize, they will have the zero value according to the type.
	// e.g. string is "", int is 0
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}


