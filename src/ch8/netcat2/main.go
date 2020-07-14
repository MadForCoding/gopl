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

	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(writer io.Writer, reader io.Reader){
	if _, err := io.Copy(writer, reader); err != nil{
		log.Fatal(err)
	}
}
