// Nonempty is an example of an in-place slice algorithm
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonEmpty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
}

func nonEmpty(str []string) []string {
	var i = 0

	for _, elem := range str {
		if elem != "" {
			str[i] = elem
			i++
		}
	}

	return str[:i]
}

func remove(str []string, index int) []string {
	copy(str[index:], str[index+1:])
	return str[:len(str)-1]
}
