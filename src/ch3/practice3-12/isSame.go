package main

import "fmt"

func main() {
	fmt.Println(isSame("chen", "ench"))
}

func isSame(s1, s2 string) bool {
	hashmap := make(map[byte]int)
	// utf-8 foreach
	for i := 0 ; i < len(s1); i++ {
		hashmap[s1[i]]++
	}

	// unicode foreach
	//hash := make(map[rune]int)
	//for _,v := range s1 {
	//	hash[v]++
	//}

	for i := 0 ; i < len(s2); i++ {
		hashmap[s2[i]]--
	}

	for _, v := range hashmap{
		if v != 0 {
			return false
		}
	}

	return true
}
