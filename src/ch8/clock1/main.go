package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main(){
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}

	for{
		// 该方法会阻塞等待连接的到来
		conn, err := listener.Accept()
		if err != nil{
			log.Print(err)
			continue
		}

		handleConn(conn)
	}

}

func handleConn(conn net.Conn){
	defer conn.Close()

	for{
		_, err := io.WriteString(conn, time.Now().Format("03:04:05PM\n"))
		if err != nil{
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
