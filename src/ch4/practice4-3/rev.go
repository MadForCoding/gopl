package main

import "fmt"

func main() {
	arr := [5]int{1,2,3,4,5}
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(arr *[5]int) {
	for left, right := 0, len(arr)-1; left < right; left,right = left+1, right-1{
		arr[left], arr[right] = arr[right], arr[left]
	}

}
