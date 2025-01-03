package main

import (
	"fmt"
	"net"
)

func serve_redis(c chan int) {
	fmt.Println("starting up redis protocol server at port 6379.")
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		c <- 1
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		c <- 1
	}

	fmt.Printf("new connection on %v\n", conn.RemoteAddr())
	defer conn.Close()

	fmt.Println("listening for redis commands...")
	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// ignore request, send back pong
		conn.Write([]byte("+OK\r\n"))
	}

}
