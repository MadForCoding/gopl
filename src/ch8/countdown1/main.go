package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown")
	tick := time.Tick(1 * time.Second)

	for count:=10; count > 0; count--{
		fmt.Println(count)
		<-tick
	}
	launch()
}

func launch(){
	fmt.Println("Lift off!!")
}
