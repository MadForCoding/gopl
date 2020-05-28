// Cf converts its numeric arguement to Celsius and Fahrenheit
package main

import (
	"ch2/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		rawTemp, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(rawTemp)
		c := tempconv.Celsius(rawTemp)
		k := tempconv.Kelvin(rawTemp)

		fmt.Printf("%s = %s, %s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KtoC(k))

	}
}
