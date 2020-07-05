package main

import "fmt"

func main(){
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"



	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
}

func sum(x ...int) int{
	var res int
	for _, val := range x{
		res +=val
	}
	return res
}

