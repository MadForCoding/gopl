package main

import "fmt"

func main() {
	str := []string{"chen", "chen", "chen", "cc", "lin", "lin", "rr", "rr"}
	str = dupRemove(str)
	fmt.Println(str)
}

func dupRemove(str []string) []string{
	var flag = true
	for flag {
		strLen := len(str) - 1
		flag = false
		for i := 0; i < strLen; i++ {
			if str[i] == str[i+1]{
				str = append(str[:i], str[i+1:]...)
				flag = true
				break
			}
		}
	}

	return str

}
