package main

import (
	"fmt"
	"time"
)

func main(){
	go spinner(1000* time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\nFibonacci(%d) = %d", n, fibN)
}

func spinner(delay time.Duration){
	for{
		for _, elem := range `-\|/`{
			fmt.Printf("\r%c", elem)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int{
	if x < 2{
		return x
	}

	return fib(x-1)+fib(x-2)
}


