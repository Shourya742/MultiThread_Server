package main

import (
	"log"
	"net"
	"time"
)
func do(conn net.Conn){
	buf := make([]byte,1024)
	_,err := conn.Read(buf)
	if err!=nil {
		log.Fatal(err)
	}
	log.Println("Processing the request")
	time.Sleep(1*time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n Hello, World!\r\n"))
	conn.Close()
}

func main(){
	Listener,err := net.Listen("tcp",":1729")
	if err!=nil {
		log.Fatal(err)
	}
	
	for{
		log.Println("Waiting for a client to connect")
		conn,err := Listener.Accept()
		if err!=nil {
			log.Fatal(err)
		}
		log.Println("client Connected")
		go do(conn)
	}
	
}