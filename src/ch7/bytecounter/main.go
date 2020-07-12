package main

import "fmt"

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error){
	*b += ByteCounter(len(p))

	return len(p), nil
}

func main(){
	var c ByteCounter
	c.Write([]byte("HELLO"))
	fmt.Println(c)

	// reset the counter
	c = 0
	name := "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)


}

