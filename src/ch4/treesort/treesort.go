package main

import "fmt"

type node struct {
	value int
	left, right *node
}




func main() {
	var slice = []int{4,3,6,1,3,8,3,4,9,23,22,67}

	sort(slice)
	fmt.Println(slice)
}

func sort(slice []int) {
	var root *node
	for _, elem := range slice {
		root = add(root, elem)
	}

	output(slice[:0], root)
}

func add(head *node, value int) *node{
	if head == nil {
		point := new(node)
		point.value = value
		head = point
	} else if value > head.value {
		head.right = add(head.right, value)
	} else {
		head.left = add(head.left, value)
	}

	return head
}

func output(slice []int, head *node) []int {
	if head != nil {
		slice = output(slice, head.left)
		slice = append(slice, head.value)
		slice = output(slice, head.right)
	}

	return slice
}
