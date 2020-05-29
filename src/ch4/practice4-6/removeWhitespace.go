package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := []byte("chen   ccc   aaa")
	str = remove(str)
	fmt.Println(string(str))

}

func remove(str []byte) []byte{
	var flag = true
	for flag {
		strLen := len(str) - 1
		flag = false
		for i := 0; i< strLen; i++ {
			if unicode.IsSpace(rune(str[i])) && unicode.IsSpace(rune(str[i+1])) {
				str = append(str[:i],str[i+1:]...)
				flag = true
				break
			}
		}
	}

	return str
}
