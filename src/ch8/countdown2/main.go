package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown")
	tick := time.Tick(1 * time.Second)

	abort := make(chan struct{})

	go func(){
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}

	}()

	for count:=10; count > 0; count--{

		select {
		case <- tick:
			fmt.Println(count)
		case <- abort:
			log.Println("Abort to launch")
			return
		}
	}
	launch()
}

func launch(){
	fmt.Println("Lift off!!")
}
