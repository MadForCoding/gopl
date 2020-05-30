package main

import (
	"fmt"
	"sort"
)

func main() {

	hashmap := map[string]int{
		"chen": 1,
		"sss": 3,
		"aaa": 0,
	}

	names := make([]string,0,len(hashmap))
	for key,_ := range hashmap {
		names = append(names,key)
	}

	sort.Strings(names)

	for _, key := range names {
		fmt.Printf("%s: %d\n", key, hashmap[key])
	}


}
