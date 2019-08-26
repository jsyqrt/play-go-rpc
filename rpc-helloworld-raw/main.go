package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func Serve() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}

func Request() {

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing error:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello, it's mayuko", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

func Demo() {
	go Serve()
	time.Sleep(1 * time.Second)
	Request()
}

func main() {
	Demo()
}
