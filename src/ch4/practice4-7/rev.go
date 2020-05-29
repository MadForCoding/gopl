package main

import "fmt"

func main() {
	var str = []byte("ccciii")
	reverse(str)
	fmt.Println(string(str))
}

func reverse(str []byte) {
	for left, right :=0,len(str)-1; left<right; left,right = left+1,right-1{
		str[left],str[right] = str[right],str[left]
	}
}
