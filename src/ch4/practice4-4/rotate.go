package main

import "fmt"

func main() {
	sli := []int{1,2,3,4,5}
	rotate(sli,2)
	fmt.Println(sli)
}

func rotate(sli []int, index int) {
	reverse(sli)
	reverse(sli[:index])
	reverse(sli[index:])

}

func reverse(sli []int) {
	for left, right:=0, len(sli) -1; left<right; left, right=left+1,right-1{
		sli[left],sli[right] = sli[right], sli[left]
	}
}
