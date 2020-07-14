package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	done := make(chan struct{})
	go func(){
		mustCopy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	<-done
}

func mustCopy(writer io.Writer, reader io.Reader){
	if _, err := io.Copy(writer, reader); err != nil{
		log.Fatal(err)
	}
}
