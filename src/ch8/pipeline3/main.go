package main

import "fmt"

func main(){
	naturals := make(chan int)
	squ := make(chan int)

	go counter(naturals)
	go squares(naturals, squ)
	printer(squ)

}


func counter(out  chan <- int ){
	for i:=0; i < 100; i++{
		out <- i
	}
	close(out)
}

func squares(in <- chan int, out chan <- int){
	for x := range in{
		out <- x
	}
	close(out)
}

func printer(in <- chan int){
	for x := range in{
		fmt.Println(x)
	}
}
