package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (s StringSlice) Len() int{
	return len(s)
}

func (s StringSlice) Less(i, j int) bool{
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

func main(){
	var strSlice = []string{"a", "bob", "chen", "zz", "aa"}
	var s []string
	s = strSlice
	s = append(s, "z")
	fmt.Printf("%p\n", s)
	fmt.Printf("%p", strSlice)
	fmt.Println(s)
	sort.Sort(StringSlice(strSlice))

	fmt.Println(strSlice)

	fmt.Printf("%T", s)

}
